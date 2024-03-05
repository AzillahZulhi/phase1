CREATE database UniversityCourseEnrollmentSystem

CREATE TABLE Student (
    student_id INT PRIMARY KEY,
    student_name VARCHAR(100) NOT NULL
);

CREATE TABLE Professor (
    professor_id INT PRIMARY KEY,
    professor_name VARCHAR(100) NOT NULL
);

CREATE TABLE Course (
    course_id INT PRIMARY KEY,
    course_title VARCHAR(200),
    max_capacity INT
);

CREATE TABLE Enrollment (
    enrollment_id INT PRIMARY KEY,
    student_id INT,
    course_id INT,
    enrollment_date DATE,
    FOREIGN KEY (student_id) REFERENCES Student(student_id),
    FOREIGN KEY (course_id) REFERENCES Course(course_id)
);

CREATE TABLE TeachingAssignment (
    assignment_id INT PRIMARY KEY,
    professor_id INT,
    course_id INT,
    start_date DATE,
    FOREIGN KEY (professor_id) REFERENCES Professor(professor_id),
    FOREIGN KEY (course_id) REFERENCES Course(course_id)
);

INSERT INTO Student (student_id, student_name)
VALUES
(1, 'John Doe'),
(2, 'Jane Smith'),
(3, 'Alice Johnson');

INSERT INTO Professor (professor_id, professor_name)
VALUES
(1, 'Dr. Smith'),
(2, 'Prof. Johnson'),
(3, 'Dr. Lee');

INSERT INTO Course (course_id, course_title, max_capacity)
VALUES
(101, 'Introduction to Programming', 30),
(102, 'Electrical Circuit Analysis', 25),
(103, 'Cell Biology', 35);

INSERT INTO Enrollment (enrollment_id, student_id, course_id, enrollment_date)
VALUES
(1, 1, 101, '2024-02-15'),
(2, 2, 102, '2024-02-17'),
(3, 3, 103, '2024-02-20');

INSERT INTO TeachingAssignment (assignment_id, professor_id, course_id, start_date)
VALUES
(1, 1, 101, '2024-02-10'),
(2, 2, 102, '2024-02-12'),
(3, 3, 103, '2024-02-14');

SELECT s.student_name
FROM Student s
JOIN Enrollment e ON s.student_id = e.student_id
WHERE e.course_id = (SELECT course_id FROM Course WHERE course_title = 'Cell Biology');

SELECT c.course_title
FROM Course c
JOIN TeachingAssignment ta ON c.course_id = ta.course_id
WHERE ta.professor_id = (SELECT professor_id FROM Professor WHERE professor_name = 'Prof. Johnson');

SELECT p.professor_name
FROM Professor p
JOIN TeachingAssignment ta ON p.professor_id = ta.professor_id
WHERE ta.course_id = (SELECT course_id FROM Course WHERE course_title = 'Electrical Circuit Analysis');