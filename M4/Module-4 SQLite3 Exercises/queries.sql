-- Q1: List employees hired after January 1, 2021
SELECT 
    first_name || ' ' || last_name as employee_name,
    hire_date
FROM Employees
WHERE hire_date > '2021-01-01'
ORDER BY hire_date;

-- Q2: Average salary by department
SELECT 
    d.dept_name,
    ROUND(AVG(e.salary), 2) as average_salary
FROM Departments d
LEFT JOIN Employees e ON d.dept_id = e.dept_id
GROUP BY d.dept_name
HAVING average_salary IS NOT NULL
ORDER BY average_salary DESC;

-- Q3: Department with highest total salary
SELECT 
    d.dept_name,
    SUM(e.salary) as total_salary
FROM Departments d
JOIN Employees e ON d.dept_id = e.dept_id
GROUP BY d.dept_name
ORDER BY total_salary DESC
LIMIT 1;

-- Q4: Departments with no employees
SELECT 
    dept_name,
    location
FROM Departments d
WHERE NOT EXISTS (
    SELECT 1 
    FROM Employees e 
    WHERE e.dept_id = d.dept_id
);

-- Q5: Employee details with department names
SELECT 
    e.first_name || ' ' || e.last_name as employee_name,
    e.hire_date,
    e.salary,
    d.dept_name,
    d.location
FROM Employees e
JOIN Departments d ON e.dept_id = d.dept_id
ORDER BY d.dept_name, employee_name;