CREATE database StudentCourseRegistrationSystem

CREATE TABLE students (
    student_id INT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    major VARCHAR(255) NOT NULL,
    year_of_study INT
);

CREATE TABLE courses (
    course_id INT PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    instructor VARCHAR(255) NOT NULL,
    schedule VARCHAR(255) NOT NULL,
    credit_hours INT
);

CREATE TABLE student_courses (
    student_id INT,
    course_id INT,
    preferred_schedule VARCHAR(255),
    PRIMARY KEY (student_id, course_id),
    FOREIGN KEY (student_id) REFERENCES students(student_id),
    FOREIGN KEY (course_id) REFERENCES courses(course_id)
);

INSERT INTO students (student_id, name, email, major, year_of_study)
VALUES
(1, 'Aboy Boiboy', 'boboiboy@example.com', 'Computer Science', 2),
(2, 'Luke ', 'luke@example.com', 'Electrical Engineering', 3),
(3, 'Slebew', 'slebew@example.com', 'Biology', 1),
(4, 'Budiman', 'budiman@example.com', 'Psychology', 4),
(5, 'Bob Marley', 'bobmarley@example.com', 'Mathematics', 2);

INSERT INTO courses (course_id, title, instructor, schedule, credit_hours)
VALUES
(101, 'Introduction to Programming', 'Dr. Smith', 'Mon-Wed-Fri 10:00 AM - 11:30 AM', 3),
(102, 'Electrical Circuit Analysis', 'Prof. Johnson', 'Tue-Thu 1:00 PM - 2:30 PM', 4),
(103, 'Cell Biology', 'Dr. Lee', 'Mon-Wed-Fri 2:00 PM - 3:30 PM', 3),
(104, 'Introduction to Psychology', 'Prof. Adams', 'Tue-Thu 10:00 AM - 11:30 AM', 3),
(105, 'Calculus I', 'Dr. Brown', 'Mon-Wed-Fri 8:00 AM - 9:30 AM', 4);

INSERT INTO student_courses (student_id, course_id, preferred_schedule)
VALUES
(1, 101, 'Mon-Wed-Fri 10:00 AM - 11:30 AM'),
(2, 102, 'Tue-Thu 1:00 PM - 2:30 PM'),
(3, 103, 'Mon-Wed-Fri 2:00 PM - 3:30 PM'),
(4, 104, 'Tue-Thu 10:00 AM - 11:30 AM'),
(5, 105, 'Mon-Wed-Fri 8:00 AM - 9:30 AM');

SELECT students.*
FROM students
JOIN student_courses ON students.student_id = student_courses.student_id
WHERE student_courses.course_id = (SELECT course_id FROM courses WHERE title = 'Introduction to Programming');

SELECT courses.*
FROM courses
JOIN student_courses ON courses.course_id = student_courses.course_id
WHERE student_courses.student_id = (SELECT student_id FROM students WHERE name = 'Bob Marley');

SELECT preferred_schedule
FROM student_courses
WHERE student_id = (SELECT student_id FROM students WHERE name = 'John Doe')
AND course_id = (SELECT course_id FROM courses WHERE title = 'Calculus I');