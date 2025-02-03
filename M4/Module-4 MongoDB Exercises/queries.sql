db.orders.aggregate([
  {
    $lookup: {
      from: "users",
      localField: "userId",
      foreignField: "userId",
      as: "userDetails"
    }
  },
  { $unwind: "$userDetails" },
  {
    $group: {
      _id: "$userId",
      totalSpent: { $sum: "$totalAmount" },
      userName: { $first: "$userDetails.name" },
      email: { $first: "$userDetails.email" }
    }
  },
  {
    $match: {
      totalSpent: { $gt: 500 }
    }
  },
  {
    $project: {
      _id: 0,
      userId: "$_id",
      userName: 1,
      email: 1,
      totalSpent: 1
    }
  }
]);

db.products.aggregate([
  { $unwind: "$ratings" },
  {
    $group: {
      _id: "$productId",
      productName: { $first: "$name" },
      averageRating: { $avg: "$ratings.rating" },
      numberOfRatings: { $sum: 1 }
    }
  },
  {
    $match: {
      averageRating: { $gte: 4 }
    }
  },
  {
    $sort: {
      averageRating: -1
    }
  },
  {
    $project: {
      _id: 0,
      productId: "$_id",
      productName: 1,
      averageRating: 1,
      numberOfRatings: 1
    }
  }
]);

db.orders.aggregate([
  {
    $match: {
      orderDate: {
        $gte: ISODate("2024-12-01T00:00:00Z"),
        $lte: ISODate("2024-12-31T23:59:59Z")
      }
    }
  },
  {
    $lookup: {
      from: "users",
      localField: "userId",
      foreignField: "userId",
      as: "userDetails"
    }
  },
  { $unwind: "$userDetails" },
  {
    $project: {
      orderId: 1,
      orderDate: 1,
      totalAmount: 1,
      status: 1,
      "userName": "$userDetails.name",
      _id: 0
    }
  },
  {
    $sort: {
      orderDate: 1
    }
  }
]);

const updateStockForOrder = async (orderId) => {
  const session = await db.startSession();
  session.startTransaction();
  
  try {
    const order = await db.orders.findOne({ orderId: orderId });
    
    for (const item of order.items) {
      await db.products.updateOne(
        { productId: item.productId },
        { $inc: { stock: -item.quantity } },
        { session }
      );
      
      const updatedProduct = await db.products.findOne(
        { productId: item.productId },
        { session }
      );
      
      if (updatedProduct.stock < 0) {
        throw new Error(`Insufficient stock for product ${item.productId}`);
      }
    }
    
    await session.commitTransaction();
  } catch (error) {
    await session.abortTransaction();
    throw error;
  } finally {
    session.endSession();
  }
};

db.warehouses.aggregate([
  {
    $geoNear: {
      near: {
        type: "Point",
        coordinates: [-74.006, 40.7128]
      },
      distanceField: "distance",
      maxDistance: 50000,
      spherical: true,
      query: {
        products: "P001"
      }
    }
  },
  {
    $project: {
      warehouseId: 1,
      distance: 1,
      products: 1,
      _id: 0
    }
  },
  {
    $sort: {
      distance: 1
    }
  }
]);