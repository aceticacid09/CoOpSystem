// frontend/src/data/users.js

export const users = {
  student: [
    {
      id: 'S001',
      username: 'student1',
      password: '123',
      email: 'akkrawit.jan@email.com',
      name: 'อัครวิทย์ จันทรัง',
      studentId: '6410110001',
      department: 'คอมพิวเตอร์',
      year: 4,
      avatar: 'https://i.pinimg.com/736x/4e/38/e7/4e38e73208c8a9c2410e4f1d9cb90ee5.jpg'
    },
    {
      id: 'S002',
      username: 'student2',
      password: '123',
      email: 'sirapop.wit@email.com',
      name: 'ศิรภพ วิทยากร',
      studentId: '6410110002',
      department: 'คณิตศาสตร์',
      year: 3,
      avatar: 'https://i.pinimg.com/736x/97/0b/4f/970b4f30501bfe2bbc06c08bee62accf.jpg'
    }
  ],

  teacher: [
    {
      id: 'T001',
      username: 'teacher1',
      password: '123',
      email: 'satjaporn.w@su.ac.th',
      name: 'ผศ.ดร.สัจอาภรณ์ ไววรรยา',
      department: 'คอมพิวเตอร์',
      position: 'ผู้ช่วยศาสตราจารย์',
      avatar: 'https://www.cp.su.ac.th/image/crop/856'
    },
    {
      id: 'T002',
      username: 'teacher2',
      password: '123',
      email: 'somchai.j@su.ac.th',
      name: 'อ.สมชาย ใจดี',
      department: 'คณิตศาสตร์',
      position: 'อาจารย์',
      avatar: 'https://i.pinimg.com/736x/63/53/d9/6353d9fff14cc31af369dd0254fd8c97.jpg'
    }
  ],

  company: [
    {
      id: 'C001',
      username: 'company1',
      password: '123',
      email: 'hr@techinnovations.com',
      name: 'บริษัท เทคโนโลยีนวัตกรรม จำกัด',
      companyType: 'เทคโนโลยีสารสนเทศ',
      address: 'เขตปทุมวัน กรุงเทพมหานคร',
      avatar: 'https://i.pinimg.com/736x/97/0b/4f/970b4f30501bfe2bbc06c08bee62accf.jpg'
    },
    {
      id: 'C002',
      username: 'company2',
      password: '123',
      email: 'contact@bioresearch.co.th',
      name: 'สถาบันวิจัยชีววิทยา',
      companyType: 'การวิจัยและพัฒนา',
      address: 'เขตจตุจักร กรุงเทพมหานคร',
      avatar: 'https://i.pinimg.com/736x/97/0b/4f/970b4f30501bfe2bbc06c08bee62accf.jpg'
    },
    {
      id: 'C003',
      username: 'company3',
      password: '123',
      email: 'info@dataanalytics.com',
      name: 'บริษัท ดาต้าอนาลิติกส์ จำกัด',
      companyType: 'วิเคราะห์ข้อมูล',
      address: 'เขตบางรัก กรุงเทพมหานคร',
      avatar: 'https://i.pinimg.com/736x/97/0b/4f/970b4f30501bfe2bbc06c08bee62accf.jpg'
    },
    {
      id: 'C004',
      username: 'company4',
      password: '123',
      email: 'hr@chemsolutions.co.th',
      name: 'บริษัท เคมีโซลูชั่นส์ จำกัด',
      companyType: 'อุตสาหกรรมเคมี',
      address: 'เขตสาทร กรุงเทพมหานคร',
      avatar: 'https://i.pinimg.com/736x/97/0b/4f/970b4f30501bfe2bbc06c08bee62accf.jpg'
    },
    {
      id: 'C005',
      username: 'company5',
      password: '123',
      email: 'contact@aifuture.com',
      name: 'บริษัท เอไอฟิวเจอร์ จำกัด',
      companyType: 'ปัญญาประดิษฐ์',
      address: 'เขตราชเทวี กรุงเทพมหานคร',
      avatar: 'https://i.pinimg.com/736x/97/0b/4f/970b4f30501bfe2bbc06c08bee62accf.jpg'
    }
  ]
};

export const authenticateUser = (username, password, role) => {
  const userList = users[role];  // ใช้ role ตรงๆ เลย
  
  if (!userList) return null;
  
  const user = userList.find(
    u => u.username === username && u.password === password
  );
  
  return user || null;
};

export const getUsersByRole = (role) => {
  return users[role] || [];
};