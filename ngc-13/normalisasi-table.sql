--------------------------------------------------------------------
| Student_ID   | Student_Name   | Course            | Course_Instructor  |
--------------------------------------------------------------------
| 101          | John           | HTML              | Mrs. Ayu           |
| 102          | Jane           | CSS               | Mrs. Amel          |
| 103          | Mary           | BasicGO           | Mr. Aldi           |
| 104          | Sarah          | HandlingError     | Mr. Safran         |
| 105          | David          | REST API          | Mr. Tugas          |
---------------------------------------------------------------------
/*
Normalisasi ke dalam 1NF
    Tabel sudah Memenuhi 1NF karena setiap kolom berisi satu nilai
Normalisasi ke dalam 2NF
    Kunci utama dapat menjadi kombinasi dari kolom Student_ID dan Course. 
    Ini karena mungkin ada beberapa mahasiswa yang mengambil beberapa kursus. 
    Oleh karena itu, tidak ada ketergantungan parsial dari atribut non-kunci ke kunci utama.
Normalisasi ke dalam 3NF
    Student_ID dan Student_Name sudah memenuhi 3NF karena tidak ada ketergantungan transitif.
    Course dan Course_Instructor juga tidak memiliki ketergantungan transitif.
*/