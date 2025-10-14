<!-- frontend/src/pages/Student/StudentProfile.vue -->
<template>
  <DashboardLayout role="student">
    <section class="profile-section container">
      <!-- Tabs -->
      <div class="tabs">
        <button v-for="tab in tabs" :key="tab" class="tab-btn" :class="{ active: activeTab === tab }"
          @click="activeTab = tab">
          {{ tab }}
        </button>
      </div>

      <!-- ========== ข้อมูลนักศึกษา (Read-only) ========== -->
      <div v-if="activeTab === 'ข้อมูลนักศึกษา'" class="profile-content">
        <div class="profile-card">
          <!-- Profile Header with Avatar -->
          <div class="profile-header">
            <div class="avatar-section">
              <div class="avatar">
                <img v-if="studentData.avatarUrl" :src="studentData.avatarUrl" alt="Profile" />
                <svg v-else xmlns="http://www.w3.org/2000/svg" width="80" height="80" viewBox="0 0 24 24" fill="none"
                  stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                  <path d="M20 21v-2a4 4 0 0 0-4-4H8a4 4 0 0 0-4 4v2"></path>
                  <circle cx="12" cy="7" r="4"></circle>
                </svg>
              </div>
            </div>

            <!-- Student Info Grid -->
            <div class="info-grid">
              <!-- Row 1 -->
              <div class="info-row">
                <div class="info-group">
                  <label>ชื่อ - นามสกุล (ไทย)</label>
                  <div class="input-readonly">{{ studentData.nameThai }}</div>
                </div>
                <div class="info-group">
                  <label>ชื่อ - นามสกุล (อังกฤษ)</label>
                  <div class="input-readonly">{{ studentData.nameEnglish }}</div>
                </div>
              </div>

              <!-- Row 2 -->
              <div class="info-row">
                <div class="info-group">
                  <label>ภาควิชา</label>
                  <div class="input-readonly">{{ studentData.department }}</div>
                </div>
                <div class="info-group">
                  <label>ภาควิชา/สาขา</label>
                  <div class="input-readonly">{{ studentData.major }}</div>
                </div>
              </div>

              <!-- Row 3 -->
              <div class="info-row">
                <div class="info-group">
                  <label>รหัสนักศึกษา</label>
                  <div class="input-readonly">{{ studentData.studentId }}</div>
                </div>
                <div class="info-group">
                  <label>หลักสูตร</label>
                  <div class="input-readonly">{{ studentData.program }}</div>
                </div>
              </div>

              <!-- Row 4 -->
              <div class="info-row">
                <div class="info-group full-width">
                  <label>อีเมล</label>
                  <div class="input-readonly">{{ studentData.email }}</div>
                </div>
              </div>

              <!-- Row 5 -->
              <div class="info-row">
                <div class="info-group">
                  <label>ระดับการศึกษา</label>
                  <div class="input-readonly">{{ studentData.educationLevel }}</div>
                </div>
                <div class="info-group">
                  <label>ชั้นปี</label>
                  <div class="input-readonly">{{ studentData.year }}</div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- ========== ข้อมูลการสมัครงาน (Editable) ========== -->
      <div v-else-if="activeTab === 'ข้อมูลการสมัครงาน'" class="application-content">
        <!-- ข้อมูลส่วนตัว Section -->
        <div class="section-wrapper">
          <h3 class="section-title-outside">ข้อมูลส่วนตัว</h3>
          <div class="section-card">
            <div class="contact-info">
              <h4>{{ applicationData.fullName }}</h4>
              <p class="info-item">
                <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none"
                  stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                  <path d="M4 4h16c1.1 0 2 .9 2 2v12c0 1.1-.9 2-2 2H4c-1.1 0-2-.9-2-2V6c0-1.1.9-2 2-2z">
                  </path>
                  <polyline points="22,6 12,13 2,6"></polyline>
                </svg>
                {{ applicationData.email }}
              </p>
              <p class="info-item">
                <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none"
                  stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                  <path
                    d="M22 16.92v3a2 2 0 0 1-2.18 2 19.79 19.79 0 0 1-8.63-3.07 19.5 19.5 0 0 1-6-6 19.79 19.79 0 0 1-3.07-8.67A2 2 0 0 1 4.11 2h3a2 2 0 0 1 2 1.72 12.84 12.84 0 0 0 .7 2.81 2 2 0 0 1-.45 2.11L8.09 9.91a16 16 0 0 0 6 6l1.27-1.27a2 2 0 0 1 2.11-.45 12.84 12.84 0 0 0 2.81.7A2 2 0 0 1 22 16.92z">
                  </path>
                </svg>
                {{ applicationData.phone }}
              </p>
            </div>
            <button class="btn-edit" @click="openEditContactPanel">แก้ไข</button>
          </div>
        </div>

        <!-- เรซูเม่ Section -->
        <div class="section-wrapper">
          <h3 class="section-title-outside">เรซูเม่</h3>
          <div class="section-card">
            <p class="helper-text">อัปโหลดเรซูเม่เพื่อพิจารณาสมัครในโครงการสหกิจศึกษาและเข้าสู่ทีมได้ทุกที่</p>
            <button class="btn-add" @click="openUploadResumePanel">เพิ่มเรซูเม่</button>

            <div v-if="applicationData.resumes.length === 0" class="empty-state-small">
              <p>ยังไม่มีเรซูเม่</p>
            </div>

            <div v-else class="resume-list">
              <div v-for="(resume, index) in applicationData.resumes" :key="index" class="resume-item">
                <div class="resume-info">
                  <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none"
                    stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                    <path d="M14 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8z"></path>
                    <polyline points="14 2 14 8 20 8"></polyline>
                    <line x1="16" y1="13" x2="8" y2="13"></line>
                    <line x1="16" y1="17" x2="8" y2="17"></line>
                    <polyline points="10 9 9 9 8 9"></polyline>
                  </svg>
                  <span>{{ resume.name }}</span>
                </div>
                <button class="btn-delete" @click="deleteResume(index)" aria-label="ลบไฟล์">
                  <svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none"
                    stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                    <polyline points="3 6 5 6 21 6"></polyline>
                    <path d="M19 6v14a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2V6m3 0V4a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v2">
                    </path>
                  </svg>
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- ========== SIDE PANEL: แก้ไขข้อมูลส่วนตัว ========== -->
      <div v-if="showEditContactPanel" class="side-panel-overlay" @click.self="closeEditContactPanel">
        <div class="side-panel">
          <button class="side-close" @click="closeEditContactPanel">&times;</button>
          <div class="side-panel-inner">
            <h3 class="side-panel-title">แก้ไขข้อมูลส่วนตัว</h3>

            <div class="form-row">
              <div class="form-group">
                <label>ชื่อ</label>
                <input type="text" v-model="editData.firstName" placeholder="กรอกชื่อ" />
              </div>
              <div class="form-group">
                <label>นามสกุล</label>
                <input type="text" v-model="editData.lastName" placeholder="กรอกนามสกุล" />
              </div>
            </div>

            <div class="form-group">
              <label>อีเมล (ส่วนตัวหรือของคุณ)</label>
              <input type="email" v-model="editData.email" placeholder="example@gmail.com" />
            </div>

            <div class="form-group">
              <label>เบอร์โทรศัพท์</label>
              <input type="tel" v-model="editData.phone" placeholder="0123456789" />
            </div>

            <div class="panel-actions">
              <button class="btn-cancel" @click="closeEditContactPanel">ยกเลิก</button>
              <button class="btn-save" @click="saveContactInfo">บันทึก</button>
            </div>
          </div>
        </div>
      </div>

      <!-- ========== SIDE PANEL: อัปโหลดเรซูเม่ ========== -->
      <div v-if="showUploadResumePanel" class="side-panel-overlay" @click.self="closeUploadResumePanel">
        <div class="side-panel">
          <button class="side-close" @click="closeUploadResumePanel">&times;</button>
          <div class="side-panel-inner">
            <h3 class="side-panel-title">เพิ่มเรซูเม่</h3>
            <p class="helper-text">อัปโหลดเรซูเม่เพื่อพิจารณาสมัครในโครงการสหกิจศึกษาและเข้าสู่ทีมได้ทุกที่</p>

            <div class="upload-area" @click="triggerFileInput" @dragover.prevent @drop.prevent="handleFileDrop">
              <input type="file" ref="fileInput" accept=".pdf" multiple @change="handleFileSelect"
                style="display: none;" />
              <div class="upload-icon">
                <svg xmlns="http://www.w3.org/2000/svg" width="48" height="48" viewBox="0 0 24 24" fill="none"
                  stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                  <path d="M21 15v4a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-4"></path>
                  <polyline points="17 8 12 3 7 8"></polyline>
                  <line x1="12" y1="3" x2="12" y2="15"></line>
                </svg>
              </div>
              <p class="upload-text">คลิกที่นี่เพื่อออัปโหลดไฟล์หรือลากไฟล์มาวางในบริเวณนี้</p>
              <p class="upload-subtext">รองรับไฟล์ .pdf เท่านั้น</p>
              <button class="btn-upload-trigger">อัปโหลดไฟล์</button>
            </div>

            <!-- รายการไฟล์ที่เลือก -->
            <div v-if="selectedFiles.length > 0" class="selected-files-section">
              <h4>รายการเรซูเม่</h4>
              <div class="resume-list">
                <div v-for="(file, index) in selectedFiles" :key="index" class="resume-item">
                  <div class="resume-info">
                    <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none"
                      stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                      <path d="M14 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8z"></path>
                      <polyline points="14 2 14 8 20 8"></polyline>
                    </svg>
                    <span>{{ file.name }}</span>
                  </div>
                  <button class="btn-delete" @click="removeSelectedFile(index)">
                    <svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none"
                      stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                      <polyline points="3 6 5 6 21 6"></polyline>
                      <path d="M19 6v14a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2V6m3 0V4a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v2">
                      </path>
                    </svg>
                  </button>
                </div>
              </div>
            </div>

            <!-- แสดงไฟล์ที่มีอยู่แล้ว -->
            <div v-if="applicationData.resumes.length > 0" class="existing-files-section">
              <h4>ไฟล์ที่มีอยู่แล้ว</h4>
              <div class="resume-list">
                <div v-for="(resume, index) in applicationData.resumes" :key="'existing-' + index" class="resume-item">
                  <div class="resume-info">
                    <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none"
                      stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                      <path d="M14 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8z"></path>
                      <polyline points="14 2 14 8 20 8"></polyline>
                      <line x1="16" y1="13" x2="8" y2="13"></line>
                      <line x1="16" y1="17" x2="8" y2="17"></line>
                      <polyline points="10 9 9 9 8 9"></polyline>
                    </svg>
                    <span>{{ resume.name }}</span>
                  </div>
                  <button class="btn-delete" @click="deleteResumeFromPanel(index)">
                    <svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none"
                      stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                      <polyline points="3 6 5 6 21 6"></polyline>
                      <path d="M19 6v14a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2V6m3 0V4a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v2">
                      </path>
                    </svg>
                  </button>
                </div>
              </div>
            </div>

            <div class="panel-actions">
              <button class="btn-cancel" @click="closeUploadResumePanel">ยกเลิก</button>
              <button class="btn-save" @click="uploadResumes" :disabled="selectedFiles.length === 0">บันทึก</button>
            </div>
          </div>
        </div>
      </div>
    </section>
  </DashboardLayout>
</template>

<script setup>
import DashboardLayout from "../../components/DashboardLayout.vue";
import { ref, reactive } from "vue";

// Tabs
const tabs = ["ข้อมูลนักศึกษา", "ข้อมูลการสมัครงาน"];
const activeTab = ref(tabs[0]);

// Student Data (Read-only)
const studentData = reactive({
  avatarUrl: "", // ถ้ามีรูป ใส่ URL ตรงนี้
  nameThai: "สมชาย ใจดี",
  nameEnglish: "Somchai Jaidee",
  department: "คณะวิทยาศาสตร์",
  major: "ภาควิชาวิทยาการคอมพิวเตอร์",
  studentId: "64010001",
  program: "ปริญญาตรี",
  email: "somchai.j@student.su.ac.th",
  educationLevel: "ปริญญาตรี",
  year: "4"
});

// Application Data (Editable)
const applicationData = reactive({
  fullName: "สมชาย ใจดี",
  email: "somchai@gmail.com",
  phone: "0123456789",
  resumes: [
    { name: "Somchai_Resume.pdf" },
    { name: "Somchai_Resume2.pdf" }
  ]
});

// Panel states
const showEditContactPanel = ref(false);
const showUploadResumePanel = ref(false);
const selectedFiles = ref([]);
const fileInput = ref(null);

// Edit data
const editData = reactive({
  firstName: "",
  lastName: "",
  email: "",
  phone: ""
});

// Methods
const openEditContactPanel = () => {
  const names = applicationData.fullName.split(" ");
  editData.firstName = names[0] || "";
  editData.lastName = names.slice(1).join(" ") || "";
  editData.email = applicationData.email;
  editData.phone = applicationData.phone;
  showEditContactPanel.value = true;
};

const closeEditContactPanel = () => {
  showEditContactPanel.value = false;
};

const saveContactInfo = () => {
  applicationData.fullName = `${editData.firstName} ${editData.lastName}`;
  applicationData.email = editData.email;
  applicationData.phone = editData.phone;
  closeEditContactPanel();
};

const openUploadResumePanel = () => {
  selectedFiles.value = [];
  showUploadResumePanel.value = true;
};

const closeUploadResumePanel = () => {
  showUploadResumePanel.value = false;
  selectedFiles.value = [];
};

const triggerFileInput = () => {
  fileInput.value.click();
};

const handleFileSelect = (event) => {
  const files = Array.from(event.target.files);
  files.forEach(file => {
    if (file.type === "application/pdf") {
      selectedFiles.value.push(file);
    }
  });
};

const handleFileDrop = (event) => {
  const files = Array.from(event.dataTransfer.files);
  files.forEach(file => {
    if (file.type === "application/pdf") {
      selectedFiles.value.push(file);
    }
  });
};

const removeSelectedFile = (index) => {
  selectedFiles.value.splice(index, 1);
};

const uploadResumes = () => {
  selectedFiles.value.forEach(file => {
    applicationData.resumes.push({ name: file.name });
  });
  closeUploadResumePanel();
};

const deleteResume = (index) => {
  if (confirm("คุณต้องการลบเรซูเม่นี้หรือไม่?")) {
    applicationData.resumes.splice(index, 1);
  }
};

const deleteResumeFromPanel = (index) => {
  if (confirm("คุณต้องการลบเรซูเม่นี้หรือไม่?")) {
    applicationData.resumes.splice(index, 1);
  }
};
</script>

<style scoped>
@import url("https://fonts.googleapis.com/css2?family=Kanit:wght@400;500;600;700&display=swap");

* {
  font-family: "Kanit", "Prompt", "Sarabun", "Noto Sans Thai", sans-serif;
  box-sizing: border-box;
}

.profile-section {
  padding: 0;
}

/* ========== TABS ========== */
.tabs {
  display: flex;
  gap: 25px;
  border-bottom: 1px solid #eee;
  margin-bottom: 30px;
  flex-wrap: wrap;
}

.tab-btn {
  background: none;
  border: none;
  font-size: 16px;
  color: #555;
  padding-bottom: 8px;
  cursor: pointer;
  position: relative;
  transition: 0.2s;
  font-weight: 500;
}

.tab-btn.active {
  color: #037266;
  font-weight: 600;
}

.tab-btn.active::after {
  content: "";
  position: absolute;
  left: 0;
  bottom: -1px;
  width: 100%;
  height: 2px;
  background: #037266;
}

/* ========== PROFILE CONTENT (ข้อมูลนักศึกษา) ========== */
.profile-content {
  padding: 20px 0;
}

.profile-card {
  background: #fff;
  border-radius: 16px;
  padding: 40px;
  box-shadow: 0 2px 12px rgba(3, 114, 102, 0.07);
  border: 1px solid #eef0ef;
}

.profile-header {
  display: flex;
  gap: 40px;
  align-items: flex-start;
}

.avatar-section {
  flex-shrink: 0;
}

.avatar {
  width: 160px;
  height: 160px;
  border-radius: 50%;
  background: #f0f0f0;
  display: flex;
  align-items: center;
  justify-content: center;
  overflow: hidden;
  border: 3px solid #e1e1e1;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
}

.avatar img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.avatar svg {
  color: #999;
}

.info-grid {
  flex: 1;
  display: flex;
  flex-direction: column;
  gap: 20px;
}

.info-row {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 20px;
}

.info-group {
  display: flex;
  flex-direction: column;
  gap: 6px;
}

.info-group.full-width {
  grid-column: 1 / -1;
}

.info-group label {
  font-size: 14px;
  font-weight: 600;
  color: #555;
}

.input-readonly {
  padding: 12px 16px;
  background: #f5f5f5;
  border: 1px solid #e0e0e0;
  border-radius: 10px;
  font-size: 15px;
  color: #666;
  cursor: not-allowed;
}

/* ========== APPLICATION CONTENT (ข้อมูลการสมัครงาน) ========== */
.application-content {
  padding: 20px 0;
  display: flex;
  flex-direction: column;
  gap: 35px;
}

.section-wrapper {
  display: flex;
  flex-direction: column;
  gap: 16px;
}

.section-title-outside {
  font-size: 22px;
  font-weight: 700;
  color: #222;
  margin: 0;
  padding-left: 4px;
}

.section-card {
  background: #fff;
  border-radius: 16px;
  padding: 30px;
  box-shadow: 0 2px 12px rgba(3, 114, 102, 0.07);
  border: 1px solid #eef0ef;
  position: relative;
}

.contact-info {
  margin-bottom: 25px;
}

.contact-info h4 {
  font-size: 22px;
  font-weight: 700;
  color: #222;
  margin: 0 0 15px 0;
}

.info-item {
  display: flex;
  align-items: center;
  gap: 10px;
  font-size: 15px;
  color: #555;
  margin: 8px 0;
}

.info-item svg {
  color: #999;
  flex-shrink: 0;
}

.helper-text {
  font-size: 14px;
  color: #777;
  margin-bottom: 20px;
  line-height: 1.6;
}

.btn-edit,
.btn-add {
  padding: 11px 24px;
  border: none;
  border-radius: 10px;
  font-size: 15px;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.2s;
  display: inline-block;
}

.btn-edit {
  background: #f0f8f7;
  color: #037266;
  border: 2px solid #cee5e2;
}

.btn-edit:hover {
  background: #e0f2f0;
  border-color: #037266;
  transform: translateY(-1px);
  box-shadow: 0 4px 12px rgba(3, 114, 102, 0.15);
}

.btn-add {
  background: linear-gradient(135deg, #037266 0%, #026b5b 100%);
  color: #fff;
  box-shadow: 0 4px 12px rgba(3, 114, 102, 0.2);
  margin-bottom: 20px;
}

.btn-add:hover {
  background: linear-gradient(135deg, #026b5b 0%, #025a4d 100%);
  transform: translateY(-2px);
  box-shadow: 0 6px 16px rgba(3, 114, 102, 0.3);
}

.empty-state-small {
  padding: 30px;
  text-align: center;
  color: #999;
  font-size: 15px;
}

.resume-list {
  display: flex;
  flex-direction: column;
  gap: 12px;
  margin-top: 20px;
}

.resume-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 14px 18px;
  background: #f9f9f9;
  border: 1px solid #e5e5e5;
  border-radius: 10px;
  transition: all 0.2s;
}

.resume-item:hover {
  background: #f0f8f7;
  border-color: #cee5e2;
}

.resume-info {
  display: flex;
  align-items: center;
  gap: 12px;
  flex: 1;
}

.resume-info svg {
  color: #037266;
  flex-shrink: 0;
}

.resume-info span {
  font-size: 15px;
  color: #333;
  font-weight: 500;
}

.btn-delete {
  background: transparent;
  border: none;
  padding: 6px;
  cursor: pointer;
  color: #999;
  transition: all 0.2s;
  border-radius: 6px;
}

.btn-delete:hover {
  background: #fee;
  color: #e74c3c;
}

/* ========== SIDE PANEL ========== */
.side-panel-overlay {
  position: fixed;
  inset: 0;
  background: rgba(0, 0, 0, 0.4);
  display: flex;
  justify-content: flex-end;
  z-index: 1000;
}

.side-panel {
  background: #fff;
  width: 600px;
  max-width: 90%;
  height: 100%;
  padding: 40px 30px;
  overflow-y: auto;
  position: relative;
  box-shadow: -4px 0 20px rgba(0, 0, 0, 0.1);
  border-radius: 12px 0 0 12px;
  transform: translateX(100%);
  animation: slideInFromRight 0.35s forwards ease-out;
}

@keyframes slideInFromRight {
  to {
    transform: translateX(0);
  }
}

.side-close {
  position: absolute;
  top: 20px;
  right: 20px;
  width: 34px;
  height: 34px;
  border: none;
  border-radius: 50%;
  background: #f0f0f0;
  color: #666;
  font-size: 20px;
  font-weight: bold;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  box-shadow: 0 2px 6px rgba(0, 0, 0, 0.15);
  transition: background 0.2s, transform 0.2s;
}

.side-close:hover {
  background: #e0e0e0;
  transform: scale(1.05);
}

.side-panel-inner {
  padding: 46px 34px 38px 34px;
  display: flex;
  flex-direction: column;
  gap: 24px;
}

.side-panel-title {
  font-size: 24px;
  font-weight: 700;
  color: #222;
  margin: 0;
}

.form-row {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 16px;
}

.form-group {
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.form-group label {
  font-size: 14px;
  font-weight: 600;
  color: #555;
}

.form-group input {
  padding: 12px 16px;
  border: 1px solid #e0e0e0;
  border-radius: 10px;
  font-size: 15px;
  transition: all 0.2s;
  outline: none;
}

.form-group input:focus {
  border-color: #037266;
  box-shadow: 0 0 0 3px rgba(3, 114, 102, 0.1);
}

.upload-area {
  border: 2px dashed #ccc;
  border-radius: 12px;
  padding: 50px 40px;
  text-align: center;
  cursor: pointer;
  transition: all 0.3s;
  background: #fafafa;
  position: relative;
}

.upload-area:hover {
  border-color: #037266;
  background: #f0f8f7;
}

.upload-icon {
  margin-bottom: 20px;
  color: #999;
}

.upload-text {
  font-size: 15px;
  color: #333;
  font-weight: 500;
  margin: 10px 0 8px 0;
}

.upload-subtext {
  font-size: 13px;
  color: #999;
  margin: 0 0 20px 0;
}

.btn-upload-trigger {
  padding: 11px 28px;
  background: #f0f8f7;
  color: #037266;
  border: 2px solid #cee5e2;
  border-radius: 10px;
  font-size: 15px;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.2s;
  margin-top: 10px;
}

.btn-upload-trigger:hover {
  background: #e0f2f0;
  border-color: #037266;
  transform: translateY(-1px);
}

.selected-files-section,
.existing-files-section {
  margin-top: 30px;
}

.selected-files-section h4,
.existing-files-section h4 {
  font-size: 18px;
  font-weight: 600;
  color: #333;
  margin-bottom: 16px;
}

.existing-files-section {
  padding-top: 30px;
  border-top: 1px solid #eee;
}

.panel-actions {
  display: flex;
  justify-content: flex-end;
  gap: 12px;
  margin-top: 30px;
  padding-top: 25px;
  border-top: 1px solid #eee;
}

.btn-cancel,
.btn-save {
  padding: 12px 28px;
  border: none;
  border-radius: 10px;
  font-size: 15px;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.2s;
}

.btn-cancel {
  background: #f0f0f0;
  color: #666;
}

.btn-cancel:hover {
  background: #e0e0e0;
  color: #333;
}

.btn-save {
  background: linear-gradient(135deg, #037266 0%, #026b5b 100%);
  color: #fff;
  box-shadow: 0 4px 12px rgba(3, 114, 102, 0.2);
}

.btn-save:hover:not(:disabled) {
  background: linear-gradient(135deg, #026b5b 0%, #025a4d 100%);
  transform: translateY(-1px);
  box-shadow: 0 6px 16px rgba(3, 114, 102, 0.3);
}

.btn-save:disabled {
  background: #ccc;
  cursor: not-allowed;
  box-shadow: none;
  opacity: 0.6;
}

/* ========== RESPONSIVE ========== */
@media (max-width: 900px) {
  .profile-header {
    flex-direction: column;
    align-items: center;
    text-align: center;
  }

  .avatar {
    width: 140px;
    height: 140px;
  }

  .info-row {
    grid-template-columns: 1fr;
  }

  .side-panel {
    width: 100%;
    max-width: 100%;
    border-radius: 0;
  }

  .side-panel-inner {
    padding: 30px 20px;
  }
}

@media (max-width: 650px) {
  .profile-card {
    padding: 25px 20px;
  }

  .section-card {
    padding: 20px;
  }

  .section-title-outside {
    font-size: 20px;
  }

  .profile-header {
    gap: 25px;
  }

  .avatar {
    width: 120px;
    height: 120px;
  }

  .form-row {
    grid-template-columns: 1fr;
  }

  .upload-area {
    padding: 40px 20px;
  }

  .panel-actions {
    flex-direction: column-reverse;
  }

  .btn-cancel,
  .btn-save {
    width: 100%;
  }

  .side-close {
    top: 12px;
    right: 12px;
    width: 32px;
    height: 32px;
    font-size: 18px;
  }

  .side-panel-inner {
    padding: 20px 16px;
  }

  .side-panel-title {
    font-size: 20px;
  }
}
</style>