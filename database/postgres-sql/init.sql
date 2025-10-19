CREATE type role_type as ENUM ('admin', 'teacher', 'student', 'company');
CREATE type company_status as ENUM ('active', 'inactive', 'pending');
CREATE type announcement_status as ENUM ('draft', 'published', 'archived');
CREATE type job_status as ENUM ('open', 'closed', 'pending');
CREATE type application_status as ENUM ('pending', 'approved', 'rejected');
CREATE type doc_types as ENUM ('resume', 'cv', 'portfolio', 'transcript', 'other');
CREATE type evaluator_types as ENUM ('teacher', 'company');

-- สร้างตาราง User (ผู้ใช้ระบบ)
CREATE TABLE Users (
    user_id SERIAL PRIMARY KEY,
    username VARCHAR(255) NOT NULL,
    password TEXT NOT NULL,
    role role_type NOT NULL
);

-- สร้างตาราง Department (ภาค)
CREATE TABLE Department (
    dept_id SERIAL PRIMARY KEY,
    dept_name VARCHAR(255) NOT NULL
);

-- สร้างตาราง Teacher (อาจารย์)
CREATE TABLE Teacher (
    teacher_id INTEGER PRIMARY KEY REFERENCES Users(user_id),
    name_th VARCHAR(255),
    name_en VARCHAR(255),
    dept_id INTEGER REFERENCES Department(dept_id),
    email_contact VARCHAR(255),
    profile_image TEXT
);

-- สร้างตาราง Student (นักศึกษา)
CREATE TABLE Student (
    student_id INTEGER PRIMARY KEY REFERENCES Users(user_id),
    name_th VARCHAR(255),
    name_en VARCHAR(255),
    student_code INTEGER NOT NULL,
    dept_id INTEGER REFERENCES Department(dept_id),
    year INTEGER,
    education_level VARCHAR(100),
    profile_image TEXT
);

-- สร้างตาราง Company (บริษัท/หน่วยงาน)
CREATE TABLE Company (
    company_id INTEGER PRIMARY KEY REFERENCES Users(user_id),
    company_name VARCHAR(255) NOT NULL,
    status company_status DEFAULT 'pending',
    profile_image TEXT
);

-- สร้างตาราง Job_Position (ตำแหน่งงาน)
CREATE TABLE Job_Position (
    job_id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    company_id INTEGER REFERENCES Company(company_id),
    attachment_id INTEGER,
    status job_status DEFAULT 'open',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- สร้างตาราง Application (ใบสมัคร)
CREATE TABLE Application (
    application_id SERIAL PRIMARY KEY,
    job_id INTEGER REFERENCES Job_Position(job_id),
    student_id INTEGER REFERENCES Student(student_id),
    sdoc_id INTEGER REFERENCES StudentDoc(sdoc_id),
    status application_status DEFAULT 'pending',
    applied_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- สร้างตาราง File (ไฟล์)
CREATE TABLE File (
    file_id SERIAL PRIMARY KEY,
    storage_path TEXT NOT NULL,
    file_name VARCHAR(255) NOT NULL,
    file_type VARCHAR(100),
    uploaded_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- สร้างตาราง Attachment (ไฟล์แนบ)
CREATE TABLE Attachment (
    attachment_id SERIAL PRIMARY KEY,
    file_count INTEGER DEFAULT 0 NOT NULL
);

-- สร้างตาราง Attachment_File (ตารางเชื่อมโยงระหว่าง Attachment และ File)
CREATE TABLE Attachment_File (
    attachment_id INTEGER REFERENCES Attachment(attachment_id) ON DELETE CASCADE,
    file_id INTEGER REFERENCES File(file_id) ON DELETE CASCADE,
    PRIMARY KEY (attachment_id, file_id)
);

-- สร้างตาราง Announ_News (ประกาศข่าวสาร)
CREATE TABLE Announ_News (
    post_id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    attachment_id INTEGER REFERENCES Attachment(attachment_id),
    status announcement_status NOT NULL DEFAULT 'draft',
    teacher_id INTEGER REFERENCES Teacher(teacher_id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- สร้างตาราง StudentDoc (เอกสารนักศึกษา)
CREATE TABLE StudentDoc (
    sdoc_id SERIAL PRIMARY KEY,
    student_id INTEGER REFERENCES Student(student_id),
    doc_type doc_types NOT NULL,
    attachment_id INTEGER REFERENCES Attachment(attachment_id),
    uploaded_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- สร้างตาราง PublicDoc (เอกสารสาธารณะ)
CREATE TABLE PublicDoc (
    pdoc_id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    file_path TEXT NOT NULL,
    teacher_id INTEGER REFERENCES Teacher(teacher_id),
    uploaded_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- สร้างตาราง Evaluator_T (อาจารย์ผู้ประเมิน)
CREATE TABLE Evaluator_T (
    evaluator_id SERIAL PRIMARY KEY,
    teacher_id INTEGER REFERENCES Teacher(teacher_id)
);

-- สร้างตาราง Evaluator_C (บริษัทผู้ประเมิน)
CREATE TABLE Evaluator_C (
    evaluator_id SERIAL PRIMARY KEY,
    company_id INTEGER REFERENCES Company(company_id)
);

-- สร้างตาราง Evaluator (ผู้ประเมิน)
CREATE TABLE Evaluator (
    evaluator_id SERIAL PRIMARY KEY,
    evaluator_type evaluator_types NOT NULL
);

-- สร้างตาราง Evaluation (การประเมิน)
CREATE TABLE Evaluation (
    evaluation_id SERIAL PRIMARY KEY,
    round INTEGER NOT NULL,
    student_id INTEGER REFERENCES Student(student_id),
    evaluator_id INTEGER REFERENCES Evaluator(evaluator_id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- สร้างตาราง Eval_topic (หัวข้อการประเมิน)
CREATE TABLE Eval_topic (
    topic_id SERIAL PRIMARY KEY,
    evaluation_id INTEGER REFERENCES Evaluation(evaluation_id),
    title VARCHAR(255) NOT NULL
);

-- สร้างตาราง Eval_Item (รายการประเมิน)
CREATE TABLE Eval_Item (
    item_id SERIAL PRIMARY KEY,
    topic_id INTEGER REFERENCES Eval_topic(topic_id),
    criterion TEXT NOT NULL,
    score NUMERIC(5,2),
    comment TEXT
);

-- เพิ่ม Foreign Key Constraints ที่ขาดหายไป
ALTER TABLE Job_Position ADD CONSTRAINT fk_job_attachment 
    FOREIGN KEY (attachment_id) REFERENCES Attachment(attachment_id);

-- สร้าง Index เพื่อเพิ่มประสิทธิภาพ
CREATE INDEX idx_student_dept ON Student(dept_id);
CREATE INDEX idx_teacher_dept ON Teacher(dept_id);
CREATE INDEX idx_application_job ON Application(job_id);
CREATE INDEX idx_application_student ON Application(student_id);
CREATE INDEX idx_job_company ON Job_Position(company_id);
CREATE INDEX idx_evaluation_student ON Evaluation(student_id);
CREATE INDEX idx_user_role ON Users(role);


--------- ข้อมูลตัวอย่าง ---------

-- เพิ่มข้อมูลตัวอย่าง User  
INSERT INTO Users (username, password, role) VALUES  
('teacher001', 'hashed_password_2', 'teacher'),  
('teacher002', 'hashed_password_3', 'teacher'),  
('student001', 'hashed_password_4', 'student'),  
('student002', 'hashed_password_5', 'student'),  
('student003', 'hashed_password_6', 'student'),  
('company001', 'hashed_password_7', 'company'),  
('company002', 'hashed_password_8', 'company');  
  
-- เพิ่มข้อมูลตัวอย่าง Department  
INSERT INTO Department (dept_name) VALUES  
('เทคโนโลยีสารสนเทศ'),
('วิทยาการคอมพิวเตอร์');  
  
-- เพิ่มข้อมูลตัวอย่าง Teacher  
INSERT INTO Teacher (teacher_id, name_th, name_en, dept_id, email_contact, profile_image) VALUES  
(1, 'ดร.สมชาย ใจดี', 'Dr.Somchai Jaidee', 1, 'somchai@su.ac.th', 'teacher1.jpg'),  
(2, 'ผศ.สมหญิง รักเรียน', 'Asst.Prof.Somying Rakrian', 2, 'somying@su.ac.th', 'teacher2.jpg');  
  
-- เพิ่มข้อมูลตัวอย่าง Student  
INSERT INTO Student (student_id, name_th, name_en, student_code, dept_id, year, education_level, profile_image) VALUES  
(3, 'นายสมศักดิ์ เก่งมาก', 'Mr.Somsak Kengmak', 64010014, 1, 5, 'ปริญญาตรี', 'student1.jpg'),  
(4, 'นางสาวสมใส สวยงาม', 'Ms.Somsai Suayngam', 65010022, 1, 4, 'ปริญญาตรี', 'student2.jpg'),  
(5, 'นายสมปอง ขยันเรียน', 'Mr.Sompong Khayanrian', 65010026, 2, 4, 'ปริญญาตรี', 'student3.jpg');
  
-- เพิ่มข้อมูลตัวอย่าง Company  
INSERT INTO Company (company_id, company_name, status, profile_image) VALUES  
(6, 'บริษัท เทคโนโลยี จำกัด', 'active', 'company1.jpg'),  
(7, 'บริษัท ซอฟต์แวร์ โซลูชั่น จำกัด', 'active', 'company2.jpg');  
  
-- เพิ่มข้อมูลตัวอย่าง File  
INSERT INTO File (storage_path, file_name, file_type) VALUES  
('/uploads/documents/', 'job_description_1.pdf', 'pdf'),  
('/uploads/documents/', 'company_profile.pdf', 'pdf'),  
('/uploads/news/', 'announcement_1.jpg', 'jpg'),  
('/uploads/student_docs/', 'resume_student1.pdf', 'pdf'),  
('/uploads/student_docs/', 'portfolio_student1.pdf', 'pdf'),
('/uploads/news/', 'announcement_2.jpg', 'jpg'),
('/uploads/news/', 'announcement_3.pdf', 'pdf'),
('/uploads/student_docs/', 'transcript_student1.pdf', 'pdf');
  
-- เพิ่มข้อมูลตัวอย่าง Attachment  
INSERT INTO Attachment (file_count) VALUES  
(1),  -- attachment_id = 1 (for job position 1)
(1),  -- attachment_id = 2 (for job position 2)
(3),  -- attachment_id = 3 (for announcement 1, has 3 files)
(1),  -- attachment_id = 4 (for student doc 1)
(3);  -- attachment_id = 5 (for student doc 2, has 3 files)

-- เพิ่มข้อมูลตัวอย่าง Attachment_File (Junction Table)
INSERT INTO Attachment_File (attachment_id, file_id) VALUES
-- Attachment 1 (Job Position 1) has 1 file
(1, 1),
-- Attachment 2 (Job Position 2) has 1 file
(2, 2),
-- Attachment 3 (Announcement 1) has 3 files
(3, 3),
(3, 6),
(3, 7),
-- Attachment 4 (Student Doc 1) has 1 file
(4, 4),
-- Attachment 5 (Student Doc 2) has 3 files
(5, 5),
(5, 8),
(5, 4);  -- Reusing file 4 to show files can be shared
  
-- เพิ่มข้อมูลตัวอย่าง Job_Position  
INSERT INTO Job_Position (title, description, company_id, attachment_id, status) VALUES  
('นักพัฒนาซอฟต์แวร์', 'รับสมัครนักพัฒนาซอฟต์แวร์ มีประสบการณ์ 0-2 ปี', 6, 1, 'open'),  
('วิศวกรระบบ', 'รับสมัครวิศวกรระบบ สำหรับงานด้าน Infrastructure', 7, 2, 'open'),  
('นักวิเคราะห์ระบบ', 'รับสมัครนักวิเคราะห์ระบบ สำหรับโปรเจคใหม่', 7, NULL, 'open');  
  
-- เพิ่มข้อมูลตัวอย่าง Application  
INSERT INTO Application (job_id, student_id, status) VALUES  
(1, 4, 'pending'),  
(1, 5, 'approved'),  
(2, 4, 'pending'),  
(3, 3, 'rejected');  
  
-- เพิ่มข้อมูลตัวอย่าง Announ_News  
INSERT INTO Announ_News (title, description, attachment_id, teacher_id) VALUES  
('ประกาศรับสมัครงาน บริษัท ABC', 'บริษัท ABC เปิดรับสมัครตำแหน่งนักพัฒนาซอฟต์แวร์', 3, 1),  
('การสัมมนาเทคโนโลยีใหม่', 'เชิญร่วมงานสัมมนาเทคโนโลยี AI และ Machine Learning', NULL, 2);  
  
-- เพิ่มข้อมูลตัวอย่าง StudentDoc  
INSERT INTO StudentDoc (student_id, doc_type, attachment_id) VALUES  
(4, 'resume', 4),  
(4, 'portfolio', 5),  
(5, 'cv', NULL);  
  
-- เพิ่มข้อมูลตัวอย่าง PublicDoc  
INSERT INTO PublicDoc (title, file_path, teacher_id) VALUES  
('เทคนิคการสัมภาษณ์งาน', '/app/uploads/public_docs/interview_tips.pdf', 1),  
('แนวทางการเขียน Resume', '/app/uploads/public_docs/resume_guide.pdf', 2);  
  
-- เพิ่มข้อมูลตัวอย่าง Evaluator  
INSERT INTO Evaluator (evaluator_type) VALUES  
('teacher'),  
('company'),  
('teacher');  
  
-- เพิ่มข้อมูลตัวอย่าง Evaluator_T  
INSERT INTO Evaluator_T (evaluator_id, teacher_id) VALUES  
(1, 1),  
(3, 2);  
  
-- เพิ่มข้อมูลตัวอย่าง Evaluator_C  
INSERT INTO Evaluator_C (evaluator_id, company_id) VALUES  
(2, 7);  
  
-- เพิ่มข้อมูลตัวอย่าง Evaluation  
INSERT INTO Evaluation (round, student_id, evaluator_id) VALUES  
(1, 4, 1),  
(1, 4, 2),  
(2, 5, 1),  
(1, 3, 3);  
  
-- เพิ่มข้อมูลตัวอย่าง Eval_topic  
INSERT INTO Eval_topic (evaluation_id, title) VALUES  
(1, 'ทักษะการเขียนโปรแกรม'),  
(1, 'การทำงานเป็นทีม'),  
(2, 'ความรู้ทางเทคนิค'),  
(2, 'การสื่อสาร'),  
(3, 'การนำเสนอผลงาน'),  
(4, 'ความคิดสร้างสรรค์');  
  
-- เพิ่มข้อมูลตัวอย่าง Eval_Item  
INSERT INTO Eval_Item (topic_id, criterion, score, comment) VALUES  
(1, 'ความสามารถในการเขียนโค้ด', 8.5, 'เขียนโค้ดได้ดี มีความเข้าใจในหลักการ'),  
(1, 'การใช้เครื่องมือพัฒนา', 7.0, 'ควรฝึกฝนการใช้เครื่องมือเพิ่มเติม'),  
(2, 'การทำงานร่วมกับผู้อื่น', 9.0, 'ทำงานเป็นทีมได้ดีมาก'),  
(3, 'ความรู้พื้นฐาน', 8.0, 'มีความรู้พื้นฐานที่ดี'),  
(4, 'การนำเสนอ', 7.5, 'การนำเสนออยู่ในเกณฑ์ดี'),  
(5, 'ความชัดเจนในการนำเสนอ', 8.5, 'นำเสนอได้ชัดเจน เข้าใจง่าย');