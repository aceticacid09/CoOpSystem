<!-- frontend/src/pages/Student/StudentJobApplication.vue -->
<template>
  <DashboardLayout role="student">
    <div class="application-page">
      <!-- Header -->
      <div class="page-header">
        <button class="btn-back" @click="goBack">
          <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none"
            stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
            <line x1="19" y1="12" x2="5" y2="12"></line>
            <polyline points="12 19 5 12 12 5"></polyline>
          </svg>
          ย้อนกลับ
        </button>
        <h1>สมัครงาน</h1>
      </div>

      <!-- Job Info Card -->
      <div class="job-info-card">
        <img :src="jobData.logo" alt="Company Logo" class="company-logo" />
        <div class="job-info-content">
          <h2>{{ jobData.title }}</h2>
          <p class="company-name">{{ jobData.company }} • {{ jobData.department }}</p>
          <button class="btn-view-details">ดูรายละเอียดงาน</button>
        </div>
      </div>

      <!-- Progress Steps -->
      <div class="progress-steps">
        <div class="step" :class="{ active: currentStep >= 1, completed: currentStep > 1 }">
          <div class="step-number">1</div>
          <div class="step-label">เลือกเอกสาร</div>
        </div>
        <div class="step-line" :class="{ active: currentStep > 1 }"></div>
        <div class="step" :class="{ active: currentStep >= 2, completed: currentStep > 2 }">
          <div class="step-number">2</div>
          <div class="step-label">ตรวจสอบและยืนยันการสมัคร</div>
        </div>
        <div class="step-line" :class="{ active: currentStep > 2 }"></div>
        <div class="step" :class="{ active: currentStep >= 3, completed: currentStep > 3 }">
          <div class="step-number">3</div>
          <div class="step-label">ผลการสมัคร</div>
        </div>
      </div>

      <!-- Step 1: Select Documents -->
      <div v-if="currentStep === 1" class="step-content">
        <h3>ข้อมูลสมัครงาน</h3>
        <div class="section-card">

          <!-- Contact Info -->
          <div class="info-display">
            <h4>{{ applicationData.fullName }}</h4>
            <p class="info-item">
              <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none"
                stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                <path d="M4 4h16c1.1 0 2 .9 2 2v12c0 1.1-.9 2-2 2H4c-1.1 0-2-.9-2-2V6c0-1.1.9-2 2-2z"></path>
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
          <button class="btn-edit-inline" @click="editContact">แก้ไข</button>
        </div>

        <h3>เอกสารประกอบการสมัคร</h3>
        <div class="section-card">
          <p class="helper-text">เลือกเรซูเม่ที่ต้องการใช้สมัครงาน</p>

          <!-- Resume Selection -->
          <div class="document-section">
            <h4>เรซูเม่ (Resume)</h4>
            <div class="radio-group">
              <label v-for="(resume, index) in applicationData.resumes" :key="index" class="radio-item"
                :class="{ selected: selectedResume === index }">
                <input type="radio" :value="index" v-model="selectedResume" name="resume" />
                <div class="radio-content">
                  <div class="radio-circle"></div>
                  <div class="radio-label">
                    <svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none"
                      stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                      <path d="M14 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8z"></path>
                      <polyline points="14 2 14 8 20 8"></polyline>
                    </svg>
                    <span>{{ resume.name }}</span>
                  </div>
                </div>
              </label>
            </div>
            <button class="btn-upload-more" @click="uploadMoreResumes">+ อัปโหลดเอกสารใหม่</button>
          </div>

          <!-- Cover Letter Selection (Optional) -->
          <div class="document-section">
            <h4>จดหมายสมัครงาน (Cover letter) <span class="optional-badge">ไม่จำเป็น</span></h4>
            <p class="helper-text-small">อัปโหลดจดหมายสมัครงานเพื่อเพิ่มโอกาสในการผ่านการพิจารณาและเข้าสู่ทีมได้ทันที
            </p>

            <div v-if="coverLetters.length === 0" class="upload-prompt">
              <svg xmlns="http://www.w3.org/2000/svg" width="32" height="32" viewBox="0 0 24 24" fill="none"
                stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                <path d="M21 15v4a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-4"></path>
                <polyline points="17 8 12 3 7 8"></polyline>
                <line x1="12" y1="3" x2="12" y2="15"></line>
              </svg>
              <p>คลิกเพื่ออัปโหลดจดหมายสมัครงาน</p>
              <button class="btn-upload" @click="uploadCoverLetter">อัปโหลด</button>
            </div>

            <div v-else class="radio-group">
              <label v-for="(letter, index) in coverLetters" :key="index" class="radio-item"
                :class="{ selected: selectedCoverLetter === index }">
                <input type="radio" :value="index" v-model="selectedCoverLetter" name="coverLetter" />
                <div class="radio-content">
                  <div class="radio-circle"></div>
                  <div class="radio-label">
                    <svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none"
                      stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                      <path d="M14 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8z"></path>
                      <polyline points="14 2 14 8 20 8"></polyline>
                    </svg>
                    <span>{{ letter.name }}</span>
                  </div>
                  <button class="btn-remove-doc" @click.stop="removeCoverLetter(index)">
                    <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none"
                      stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                      <line x1="18" y1="6" x2="6" y2="18"></line>
                      <line x1="6" y1="6" x2="18" y2="18"></line>
                    </svg>
                  </button>
                </div>
              </label>
            </div>
          </div>

          <!-- Transcript (Optional) -->
          <div class="document-section">
            <h4>ใบแสดงผลการเรียน (Transcript) <span class="optional-badge">ไม่จำเป็น</span></h4>
            <p class="helper-text-small">อัปโหลดใบแสดงผลการเรียนเพื่อเพิ่มโอกาสในการผ่านการพิจารณาและเข้าสู่ทีมได้ทันที
            </p>
            <button class="btn-upload-transcript" @click="uploadTranscript">
              <svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none"
                stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                <path d="M21 15v4a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-4"></path>
                <polyline points="17 8 12 3 7 8"></polyline>
                <line x1="12" y1="3" x2="12" y2="15"></line>
              </svg>
              อัปโหลด
            </button>
          </div>
        </div>

        <div class="action-buttons">
          <button class="btn-secondary" @click="goBack">ยกเลิก</button>
          <button class="btn-primary" @click="nextStep" :disabled="selectedResume === null">ถัดไป</button>
        </div>
      </div>

      <!-- Step 2: Review -->
      <div v-if="currentStep === 2" class="step-content">
        <div class="section-card review-card">
          <h3>ตรวจสอบข้อมูล</h3>

          <div class="review-section">
            <h4>ข้อมูลส่วนตัว</h4>
            <div class="review-item">
              <span class="review-label">ชื่อ-นามสกุล:</span>
              <span class="review-value">{{ applicationData.fullName }}</span>
            </div>
            <div class="review-item">
              <span class="review-label">อีเมล:</span>
              <span class="review-value">{{ applicationData.email }}</span>
            </div>
            <div class="review-item">
              <span class="review-label">เบอร์โทรศัพท์:</span>
              <span class="review-value">{{ applicationData.phone }}</span>
            </div>
          </div>

          <div class="review-section">
            <h4>เอกสารที่เลือก</h4>
            <div class="review-item">
              <span class="review-label">เรซูเม่:</span>
              <span class="review-value">{{ applicationData.resumes[selectedResume]?.name }}</span>
            </div>
            <div v-if="selectedCoverLetter !== null" class="review-item">
              <span class="review-label">จดหมายสมัครงาน:</span>
              <span class="review-value">{{ coverLetters[selectedCoverLetter]?.name }}</span>
            </div>
          </div>
        </div>

        <div class="section-card confirmation-card">
          <h3>ยืนยันการสมัครงาน</h3>
          <p class="confirmation-text">คุณต้องการยืนยันการสมัครงานตำแหน่ง <strong>{{ jobData.title }}</strong> ที่บริษัท
            <strong>{{ jobData.company }}</strong> ใช่หรือไม่?</p>

          <div class="confirmation-notice">
            <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none"
              stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
              <circle cx="12" cy="12" r="10"></circle>
              <line x1="12" y1="8" x2="12" y2="12"></line>
              <line x1="12" y1="16" x2="12.01" y2="16"></line>
            </svg>
            <p>เมื่อกดยืนยันแล้ว ข้อมูลการสมัครของคุณจะถูกส่งไปยังบริษัทและไม่สามารถแก้ไขได้</p>
          </div>
        </div>

        <div class="action-buttons">
          <button class="btn-secondary" @click="prevStep">ย้อนกลับ</button>
          <button class="btn-primary" @click="nextStep">ถัดไป</button>
        </div>
      </div>

      <!-- Step 3: Application Summary -->
      <div v-if="currentStep === 3" class="step-content">
        <div class="application-summary-card">
          <!-- Header with Icon -->
          <div class="summary-header">
            <div class="summary-icon">
              <svg xmlns="http://www.w3.org/2000/svg" width="48" height="48" viewBox="0 0 24 24" fill="none"
                stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                <path d="M14 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8z"></path>
                <polyline points="14 2 14 8 20 8"></polyline>
                <line x1="16" y1="13" x2="8" y2="13"></line>
                <line x1="16" y1="17" x2="8" y2="17"></line>
                <polyline points="10 9 9 9 8 9"></polyline>
              </svg>
            </div>
            <h2>ใบสมัครงานสำเร็จเรียบร้อยแล้ว</h2>
            <p class="summary-subtitle">กรุณาตรวจสอบรายละเอียดใบสมัครของคุณ</p>
          </div>

          <!-- Job Position Section -->
          <div class="summary-section">
            <div class="section-header-summary">
              <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none"
                stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                <rect x="2" y="7" width="20" height="14" rx="2" ry="2"></rect>
                <path d="M16 21V5a2 2 0 0 0-2-2h-4a2 2 0 0 0-2 2v16"></path>
              </svg>
              <h4>ตำแหน่งงานที่สมัคร</h4>
            </div>
            <div class="summary-content">
              <div class="job-position-display">
                <img :src="jobData.logo" alt="Company Logo" class="job-logo-small" />
                <div class="job-details-summary">
                  <h5>{{ jobData.title }}</h5>
                  <p class="company-detail">{{ jobData.company }}</p>
                  <span class="department-badge">{{ jobData.department }}</span>
                </div>
              </div>
            </div>
          </div>

          <!-- Applicant Info Section -->
          <div class="summary-section">
            <div class="section-header-summary">
              <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none"
                stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                <path d="M20 21v-2a4 4 0 0 0-4-4H8a4 4 0 0 0-4 4v2"></path>
                <circle cx="12" cy="7" r="4"></circle>
              </svg>
              <h4>ข้อมูลผู้สมัคร</h4>
            </div>
            <div class="summary-content">
              <div class="applicant-info-grid">
                <div class="info-row-summary">
                  <span class="info-label-summary">ชื่อ-นามสกุล</span>
                  <span class="info-value-summary">{{ applicationData.fullName }}</span>
                </div>
                <div class="info-row-summary">
                  <span class="info-label-summary">อีเมล</span>
                  <span class="info-value-summary">
                    <svg xmlns="http://www.w3.org/2000/svg" width="14" height="14" viewBox="0 0 24 24" fill="none"
                      stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                      <path d="M4 4h16c1.1 0 2 .9 2 2v12c0 1.1-.9 2-2 2H4c-1.1 0-2-.9-2-2V6c0-1.1.9-2 2-2z"></path>
                      <polyline points="22,6 12,13 2,6"></polyline>
                    </svg>
                    {{ applicationData.email }}
                  </span>
                </div>
                <div class="info-row-summary">
                  <span class="info-label-summary">เบอร์โทรศัพท์</span>
                  <span class="info-value-summary">
                    <svg xmlns="http://www.w3.org/2000/svg" width="14" height="14" viewBox="0 0 24 24" fill="none"
                      stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                      <path
                        d="M22 16.92v3a2 2 0 0 1-2.18 2 19.79 19.79 0 0 1-8.63-3.07 19.5 19.5 0 0 1-6-6 19.79 19.79 0 0 1-3.07-8.67A2 2 0 0 1 4.11 2h3a2 2 0 0 1 2 1.72 12.84 12.84 0 0 0 .7 2.81 2 2 0 0 1-.45 2.11L8.09 9.91a16 16 0 0 0 6 6l1.27-1.27a2 2 0 0 1 2.11-.45 12.84 12.84 0 0 0 2.81.7A2 2 0 0 1 22 16.92z">
                      </path>
                    </svg>
                    {{ applicationData.phone }}
                  </span>
                </div>
              </div>
            </div>
          </div>

          <!-- Documents Section -->
          <div class="summary-section">
            <div class="section-header-summary">
              <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none"
                stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                <path d="M13 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V9z"></path>
                <polyline points="13 2 13 9 20 9"></polyline>
              </svg>
              <h4>เอกสารที่แนบมา</h4>
            </div>
            <div class="summary-content">
              <div class="documents-list">
                <!-- Resume -->
                <div class="document-item">
                  <div class="document-icon resume-icon">
                    <svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none"
                      stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                      <path d="M14 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8z"></path>
                      <polyline points="14 2 14 8 20 8"></polyline>
                    </svg>
                  </div>
                  <div class="document-info">
                    <span class="document-type">เรซูเม่ (Resume)</span>
                    <span class="document-name">{{ applicationData.resumes[selectedResume]?.name }}</span>
                  </div>
                </div>

                <!-- Cover Letter (if selected) -->
                <div v-if="selectedCoverLetter !== null" class="document-item">
                  <div class="document-icon cover-icon">
                    <svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none"
                      stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                      <path d="M14 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8z"></path>
                      <polyline points="14 2 14 8 20 8"></polyline>
                    </svg>
                  </div>
                  <div class="document-info">
                    <span class="document-type">จดหมายสมัครงาน (Cover Letter)</span>
                    <span class="document-name">{{ coverLetters[selectedCoverLetter]?.name }}</span>
                  </div>
                </div>
              </div>
            </div>
          </div>

          <!-- Application Date -->
          <div class="application-date">
            <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none"
              stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
              <rect x="3" y="4" width="18" height="18" rx="2" ry="2"></rect>
              <line x1="16" y1="2" x2="16" y2="6"></line>
              <line x1="8" y1="2" x2="8" y2="6"></line>
              <line x1="3" y1="10" x2="21" y2="10"></line>
            </svg>
            <span>สมัครเมื่อวันที่ {{ getCurrentDate() }}</span>
          </div>
        </div>

        <!-- Application Status Timeline -->
        <div class="status-timeline-card">
          <div class="timeline-header">
            <h3>ขั้นตอนการสมัครงาน</h3>
            <button class="btn-track-status">ตรวจสอบเอกสาร</button>
          </div>
          
          <div class="timeline-container">
            <div class="timeline-wrapper">
              <div v-for="(stage, index) in applicationStages" :key="index" class="timeline-item">
                <div class="timeline-dot" :class="stage.status">
                  <svg v-if="stage.status === 'passed'" xmlns="http://www.w3.org/2000/svg" width="20" height="20"
                    viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="3" stroke-linecap="round"
                    stroke-linejoin="round">
                    <polyline points="20 6 9 17 4 12"></polyline>
                  </svg>
                  <span v-else-if="stage.status === 'current'">?</span>
                  <svg v-else-if="stage.status === 'rejected'" xmlns="http://www.w3.org/2000/svg" width="20" height="20"
                    viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round"
                    stroke-linejoin="round">
                    <line x1="18" y1="6" x2="6" y2="18"></line>
                    <line x1="6" y1="6" x2="18" y2="18"></line>
                  </svg>
                </div>
                <div class="timeline-label">{{ stage.name }}</div>
                <div v-if="index < applicationStages.length - 1" class="timeline-connector"
                  :class="getConnectorClass(index)"></div>
              </div>
            </div>
          </div>

          <div class="timeline-info">
            <div class="info-box">
              <svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none"
                stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                <circle cx="12" cy="12" r="10"></circle>
                <line x1="12" y1="8" x2="12" y2="12"></line>
                <line x1="12" y1="16" x2="12.01" y2="16"></line>
              </svg>
              <p>ระบบจะแจ้งเตือนผ่านอีเมลเมื่อมีการเปลี่ยนแปลงสถานะการสมัครงานของคุณ</p>
            </div>
          </div>
        </div>

        <div class="action-buttons">
          <button class="btn-secondary" @click="prevStep">ย้อนกลับ</button>
          <button class="btn-primary btn-confirm" @click="submitApplication">ยืนยันการสมัคร</button>
        </div>
      </div>
    </div>
  </DashboardLayout>
</template>

<script setup>
import { ref, reactive, onMounted } from 'vue';
import { useRouter, useRoute } from 'vue-router';
import DashboardLayout from '../../components/DashboardLayout.vue';

const router = useRouter();
const route = useRoute();

// Get job ID from route
const jobId = ref(route.params.jobId);

// Current step
const currentStep = ref(1);

// Job Data (mock - should fetch from API)
const jobData = reactive({
  id: jobId.value,
  title: "Frontend Developer (In-tern)",
  company: "บริษัท Google จำกัด",
  department: "ภาควิชาคอมพิวเตอร์",
  logo: "https://i.pinimg.com/736x/c9/a8/bb/c9a8bbf40e4ea04310ff4cc2b94e9f95.jpg"
});

// Application Data
const applicationData = reactive({
  fullName: "สมชาย ใจดี",
  email: "somchai@gmail.com",
  phone: "0123456789",
  resumes: [
    { name: "Somchai_Resume.pdf" },
    { name: "Somchai_Resume2.pdf" }
  ]
});

// Selected documents
const selectedResume = ref(null);
const selectedCoverLetter = ref(null);
const coverLetters = ref([]);

// Application Stages for Timeline
const applicationStages = ref([
  { name: "กดสมัคร", status: "passed" },
  { name: "ตรวจสอบเอกสาร", status: "current" },
  { name: "สัมภาษณ์", status: "pending" },
  { name: "รอพิจารณา", status: "pending" },
  { name: "ประกาศผล", status: "pending" }
]);

// Methods
const goBack = () => {
  router.back();
};

const nextStep = () => {
  if (currentStep.value < 3) {
    currentStep.value++;
    window.scrollTo({ top: 0, behavior: 'smooth' });
  }
};

const prevStep = () => {
  if (currentStep.value > 1) {
    currentStep.value--;
    window.scrollTo({ top: 0, behavior: 'smooth' });
  }
};

const editContact = () => {
  // Open edit panel or redirect to profile
  alert('เปิดหน้าแก้ไขข้อมูล');
};

const uploadMoreResumes = () => {
  alert('เปิดหน้าอัปโหลดเรซูเม่');
};

const uploadCoverLetter = () => {
  // Mock upload
  coverLetters.value.push({ name: "CoverLetter_Sample.pdf" });
  selectedCoverLetter.value = 0;
};

const removeCoverLetter = (index) => {
  coverLetters.value.splice(index, 1);
  if (selectedCoverLetter.value === index) {
    selectedCoverLetter.value = null;
  }
};

const uploadTranscript = () => {
  alert('เปิดหน้าอัปโหลด Transcript');
};

const getCurrentDate = () => {
  const today = new Date();
  const day = today.getDate();
  const months = [
    "มกราคม", "กุมภาพันธ์", "มีนาคม", "เมษายน", "พฤษภาคม", "มิถุนายน",
    "กรกฎาคม", "สิงหาคม", "กันยายน", "ตุลาคม", "พฤศจิกายน", "ธันวาคม"
  ];
  const month = months[today.getMonth()];
  const year = today.getFullYear() + 543;
  return `${day} ${month} ${year}`;
};

const getConnectorClass = (index) => {
  const currentStage = applicationStages.value[index];
  const nextStage = applicationStages.value[index + 1];

  if (currentStage.status === "passed" && (nextStage.status === "passed" || nextStage.status === "current")) {
    return "passed";
  }

  if (currentStage.status === "passed" && nextStage.status === "rejected") {
    return "rejected";
  }

  return "pending";
};

const submitApplication = () => {
  // Submit application to API
  alert('สมัครงานสำเร็จ!');
  router.push({ name: 'StudentJobs' });
};

onMounted(() => {
  // Fetch job data if needed
  console.log('Job ID:', jobId.value);
});
</script>

<style scoped>
@import url("https://fonts.googleapis.com/css2?family=Kanit:wght@400;500;600;700&display=swap");

* {
  font-family: "Kanit", sans-serif;
  box-sizing: border-box;
}

.application-page {
  max-width: 900px;
  margin: 0 auto;
  padding: 20px;
}

/* ========== PAGE HEADER ========== */
.page-header {
  margin-bottom: 30px;
}

.btn-back {
  display: inline-flex;
  align-items: center;
  gap: 8px;
  padding: 10px 20px;
  background: #f0f0f0;
  border: none;
  border-radius: 10px;
  font-size: 15px;
  font-weight: 600;
  color: #666;
  cursor: pointer;
  transition: all 0.2s;
  margin-bottom: 20px;
}

.btn-back:hover {
  background: #e0e0e0;
  color: #333;
}

.page-header h1 {
  font-size: 28px;
  font-weight: 700;
  color: #222;
  margin: 0;
}

/* ========== JOB INFO CARD ========== */
.job-info-card {
  background: #fff;
  border-radius: 16px;
  padding: 25px 30px;
  box-shadow: 0 2px 12px rgba(3, 114, 102, 0.07);
  border: 1px solid #eef0ef;
  display: flex;
  align-items: center;
  gap: 25px;
  margin-bottom: 40px;
}

.company-logo {
  width: 80px;
  height: 80px;
  border-radius: 50%;
  object-fit: cover;
  border: 2px solid #e1e1e1;
  flex-shrink: 0;
}

.job-info-content {
  flex: 1;
}

.job-info-content h2 {
  font-size: 22px;
  font-weight: 700;
  color: #222;
  margin: 0 0 8px 0;
}

.company-name {
  font-size: 15px;
  color: #666;
  margin: 0 0 12px 0;
}

.btn-view-details {
  padding: 8px 18px;
  background: #f0f8f7;
  color: #037266;
  border: 1px solid #cee5e2;
  border-radius: 8px;
  font-size: 14px;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.2s;
}

.btn-view-details:hover {
  background: #e0f2f0;
  border-color: #037266;
}

/* ========== PROGRESS STEPS ========== */
.progress-steps {
  display: flex;
  align-items: center;
  justify-content: space-between;
  margin-bottom: 40px;
  padding: 0 20px;
}

.step {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 12px;
  flex: 0 0 auto;
}

.step-number {
  width: 50px;
  height: 50px;
  border-radius: 50%;
  background: #f0f0f0;
  color: #999;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 20px;
  font-weight: 700;
  transition: all 0.3s;
  border: 3px solid #f0f0f0;
}

.step.active .step-number {
  background: #037266;
  color: #fff;
  border-color: #037266;
}

.step.completed .step-number {
  background: #10b981;
  color: #fff;
  border-color: #10b981;
}

.step-label {
  font-size: 14px;
  color: #999;
  font-weight: 500;
  text-align: center;
  transition: all 0.3s;
}

.step.active .step-label {
  color: #037266;
  font-weight: 600;
}

.step-line {
  flex: 1;
  height: 3px;
  background: #f0f0f0;
  transition: all 0.3s;
  margin: 0 20px;
  margin-bottom: 35px;
}

.step-line.active {
  background: #037266;
}

/* ========== SECTION CARD ========== */
.section-card {
  background: #fff;
  border-radius: 16px;
  padding: 30px;
  box-shadow: 0 2px 12px rgba(3, 114, 102, 0.07);
  border: 1px solid #eef0ef;
  margin-bottom: 25px;
  position: relative;
}

.section-card h3 {
  font-size: 20px;
  font-weight: 700;
  color: #222;
  margin: 0 0 20px 0;
}

.section-card h4 {
  font-size: 16px;
  font-weight: 600;
  color: #333;
  margin: 0 0 12px 0;
  display: flex;
  align-items: center;
  gap: 8px;
}

.helper-text {
  font-size: 14px;
  color: #777;
  margin-bottom: 20px;
  line-height: 1.6;
}

.helper-text-small {
  font-size: 13px;
  color: #888;
  margin-bottom: 15px;
  line-height: 1.5;
}

/* ========== INFO DISPLAY ========== */
.info-display {
  margin-bottom: 20px;
}

.info-display h4 {
  font-size: 20px;
  font-weight: 700;
  color: #222;
  margin: 0 0 12px 0;
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

.btn-edit-inline {
  padding: 10px 22px;
  background: #f0f8f7;
  color: #037266;
  border: 2px solid #cee5e2;
  border-radius: 10px;
  font-size: 15px;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.2s;
}

.btn-edit-inline:hover {
  background: #e0f2f0;
  border-color: #037266;
  transform: translateY(-1px);
}

/* ========== DOCUMENT SECTION ========== */
.document-section {
  margin: 30px 0;
  padding: 25px 0;
  border-bottom: 1px solid #eee;
}

.document-section:last-child {
  border-bottom: none;
}

.optional-badge {
  font-size: 12px;
  color: #999;
  font-weight: 500;
  background: #f5f5f5;
  padding: 3px 10px;
  border-radius: 12px;
}

/* ========== RADIO GROUP ========== */
.radio-group {
  display: flex;
  flex-direction: column;
  gap: 12px;
  margin-bottom: 15px;
}

.radio-item {
  display: block;
  cursor: pointer;
}

.radio-item input[type="radio"] {
  display: none;
}

.radio-content {
  display: flex;
  align-items: center;
  gap: 15px;
  padding: 15px 18px;
  background: #fff;
  border: 2px solid #e5e5e5;
  border-radius: 12px;
  transition: all 0.2s;
}

.radio-item:hover .radio-content {
  border-color: #cee5e2;
  background: #f0f8f7;
}

.radio-item.selected .radio-content {
  border-color: #037266;
  background: #f0f8f7;
}

.radio-circle {
  width: 22px;
  height: 22px;
  border: 2px solid #d0d0d0;
  border-radius: 50%;
  position: relative;
  flex-shrink: 0;
  transition: all 0.2s;
}

.radio-item.selected .radio-circle {
  border-color: #037266;
}

.radio-item.selected .radio-circle::after {
  content: '';
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  width: 12px;
  height: 12px;
  background: #037266;
  border-radius: 50%;
}

.radio-label {
  display: flex;
  align-items: center;
  gap: 10px;
  flex: 1;
  font-size: 15px;
  font-weight: 500;
  color: #333;
}

.radio-label svg {
  color: #037266;
}

.btn-remove-doc {
  background: transparent;
  border: none;
  padding: 6px;
  cursor: pointer;
  color: #999;
  transition: all 0.2s;
  border-radius: 6px;
  display: flex;
  align-items: center;
  justify-content: center;
}

.btn-remove-doc:hover {
  background: #fee;
  color: #e74c3c;
}

/* ========== UPLOAD BUTTONS ========== */
.btn-upload-more,
.btn-upload,
.btn-upload-transcript {
  padding: 11px 24px;
  background: linear-gradient(135deg, #037266 0%, #026b5b 100%);
  color: #fff;
  border: none;
  border-radius: 10px;
  font-size: 15px;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.2s;
  display: inline-flex;
  align-items: center;
  gap: 8px;
}

.btn-upload-more:hover,
.btn-upload:hover,
.btn-upload-transcript:hover {
  background: linear-gradient(135deg, #026b5b 0%, #025a4d 100%);
  transform: translateY(-1px);
  box-shadow: 0 4px 12px rgba(3, 114, 102, 0.25);
}

.upload-prompt {
  background: #fafafa;
  border: 2px dashed #ccc;
  border-radius: 12px;
  padding: 40px 30px;
  text-align: center;
  margin-bottom: 15px;
}

.upload-prompt svg {
  color: #999;
  margin-bottom: 15px;
}

.upload-prompt p {
  font-size: 15px;
  color: #666;
  margin: 0 0 20px 0;
}

/* ========== REVIEW CARD ========== */
.review-card {
  background: #f9f9f9;
}

.review-section {
  padding: 20px 0;
  border-bottom: 1px solid #e0e0e0;
}

.review-section:last-child {
  border-bottom: none;
}

.review-section h4 {
  font-size: 18px;
  font-weight: 600;
  color: #222;
  margin-bottom: 15px;
}

.review-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 12px 0;
  gap: 20px;
}

.review-label {
  font-size: 15px;
  color: #666;
  font-weight: 500;
  min-width: 150px;
}

.review-value {
  font-size: 15px;
  color: #333;
  font-weight: 600;
  text-align: right;
  word-break: break-word;
}

/* ========== CONFIRMATION CARD ========== */
.confirmation-card {
  text-align: center;
  padding: 40px 30px;
}

.confirmation-card h3 {
  font-size: 24px;
  margin-bottom: 20px;
}

.confirmation-text {
  font-size: 16px;
  color: #555;
  line-height: 1.7;
  margin-bottom: 30px;
}

.confirmation-text strong {
  color: #037266;
  font-weight: 700;
}

.confirmation-notice {
  background: #fff9e6;
  border: 1px solid #ffd43b;
  border-radius: 12px;
  padding: 20px;
  display: flex;
  align-items: flex-start;
  gap: 15px;
  text-align: left;
}

.confirmation-notice svg {
  color: #f59e0b;
  flex-shrink: 0;
  margin-top: 2px;
}

.confirmation-notice p {
  font-size: 14px;
  color: #666;
  margin: 0;
  line-height: 1.6;
}

/* ========== APPLICATION SUMMARY CARD (STEP 3) ========== */
.application-summary-card {
  background: #f9f9f9;
  border-radius: 16px;
  padding: 40px;
  box-shadow: 0 2px 12px rgba(3, 114, 102, 0.07);
  border: 1px solid #eef0ef;
  margin-bottom: 25px;
}

.summary-header {
  text-align: center;
  margin-bottom: 40px;
  padding-bottom: 30px;
  border-bottom: 2px solid #e0e0e0;
}

.summary-icon {
  width: 80px;
  height: 80px;
  background: linear-gradient(135deg, #037266 0%, #026b5b 100%);
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  margin: 0 auto 20px;
  color: #fff;
  box-shadow: 0 4px 16px rgba(3, 114, 102, 0.2);
}

.summary-header h2 {
  font-size: 26px;
  font-weight: 700;
  color: #222;
  margin: 0 0 10px 0;
}

.summary-subtitle {
  font-size: 15px;
  color: #666;
  margin: 0;
}

.summary-section {
  background: #fff;
  border-radius: 12px;
  padding: 25px;
  margin-bottom: 20px;
  box-shadow: 0 1px 4px rgba(0, 0, 0, 0.04);
}

.section-header-summary {
  display: flex;
  align-items: center;
  gap: 10px;
  margin-bottom: 20px;
  padding-bottom: 15px;
  border-bottom: 2px solid #f0f0f0;
}

.section-header-summary svg {
  color: #037266;
  flex-shrink: 0;
}

.section-header-summary h4 {
  font-size: 18px;
  font-weight: 600;
  color: #222;
  margin: 0;
}

.summary-content {
  padding: 10px 0;
}

/* Job Position Display */
.job-position-display {
  display: flex;
  align-items: center;
  gap: 20px;
}

.job-logo-small {
  width: 60px;
  height: 60px;
  border-radius: 50%;
  object-fit: cover;
  border: 2px solid #e1e1e1;
  flex-shrink: 0;
}

.job-details-summary {
  flex: 1;
}

.job-details-summary h5 {
  font-size: 18px;
  font-weight: 700;
  color: #222;
  margin: 0 0 6px 0;
}

.company-detail {
  font-size: 14px;
  color: #666;
  margin: 0 0 8px 0;
}

.department-badge {
  display: inline-block;
  background: #e0f2f0;
  color: #037266;
  padding: 4px 12px;
  border-radius: 12px;
  font-size: 12px;
  font-weight: 600;
}

/* Applicant Info Grid */
.applicant-info-grid {
  display: flex;
  flex-direction: column;
  gap: 15px;
}

.info-row-summary {
  display: flex;
  flex-direction: column;
  gap: 6px;
  padding: 12px 0;
  border-bottom: 1px solid #f5f5f5;
}

.info-row-summary:last-child {
  border-bottom: none;
}

.info-label-summary {
  font-size: 13px;
  color: #999;
  font-weight: 500;
  text-transform: uppercase;
  letter-spacing: 0.5px;
}

.info-value-summary {
  font-size: 15px;
  color: #333;
  font-weight: 600;
  display: flex;
  align-items: center;
  gap: 8px;
}

.info-value-summary svg {
  color: #999;
  flex-shrink: 0;
}

/* Documents List */
.documents-list {
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.document-item {
  display: flex;
  align-items: center;
  gap: 15px;
  padding: 15px;
  background: #f9f9f9;
  border: 1px solid #e5e5e5;
  border-radius: 10px;
  transition: all 0.2s;
}

.document-item:hover {
  background: #f0f8f7;
  border-color: #cee5e2;
}

.document-icon {
  width: 42px;
  height: 42px;
  border-radius: 10px;
  display: flex;
  align-items: center;
  justify-content: center;
  flex-shrink: 0;
}

.resume-icon {
  background: linear-gradient(135deg, #037266 0%, #026b5b 100%);
  color: #fff;
}

.cover-icon {
  background: linear-gradient(135deg, #10b981 0%, #059669 100%);
  color: #fff;
}

.document-info {
  flex: 1;
  display: flex;
  flex-direction: column;
  gap: 4px;
}

.document-type {
  font-size: 13px;
  color: #999;
  font-weight: 500;
}

.document-name {
  font-size: 15px;
  color: #333;
  font-weight: 600;
}

/* Application Date */
.application-date {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 8px;
  padding: 15px;
  background: #fff;
  border-radius: 10px;
  font-size: 14px;
  color: #666;
  font-weight: 500;
  margin-top: 10px;
  box-shadow: 0 1px 4px rgba(0, 0, 0, 0.04);
}

.application-date svg {
  color: #037266;
  flex-shrink: 0;
}

/* ========== STATUS TIMELINE CARD (แก้ไข) ========== */
.status-timeline-card {
  background: #fff;
  border-radius: 16px;
  padding: 30px;
  box-shadow: 0 2px 12px rgba(3, 114, 102, 0.07);
  border: 1px solid #eef0ef;
  margin-top: 25px;
}

.timeline-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 30px;
  padding-bottom: 20px;
  border-bottom: 2px solid #f0f0f0;
}

.timeline-header h3 {
  font-size: 20px;
  font-weight: 700;
  color: #222;
  margin: 0;
}

.btn-track-status {
  padding: 8px 20px;
  border-radius: 20px;
  font-size: 13px;
  font-weight: 600;
  white-space: nowrap;
  transition: all 0.3s ease;
  flex-shrink: 0;
  background: #dbeafe;
  color: #1e40af;
  border: none; 
}

/* Timeline Container */
.timeline-container {
  position: relative;
  margin-bottom: 28px;
  padding: 0 10px;
}

.timeline-wrapper {
  display: flex;
  justify-content: space-between;
  align-items: flex-start;
  gap: 6px;
  position: relative;
  z-index: 2;
}

.timeline-item {
  display: flex;
  flex-direction: column;
  align-items: center;
  flex: 1;
  position: relative;
}

/* Timeline Dot - ขนาดเหมือนใน StudentHistoryJobs.vue */
.timeline-dot {
  width: 40px;
  height: 40px;
  border-radius: 50%;
  background: #f8f8f8;
  border: 3px solid #e8e8e8;
  display: flex;
  align-items: center;
  justify-content: center;
  margin-bottom: 14px;
  transition: all 0.3s ease;
  flex-shrink: 0;
  position: relative;
  z-index: 3;
  font-weight: 700;
  font-size: 20px;
  box-shadow: 0 2px 6px rgba(0, 0, 0, 0.08);
}

.timeline-dot.passed {
  background: linear-gradient(135deg, #10b981 0%, #059669 100%);
  border-color: #10b981;
  color: #fff;
  box-shadow: 0 3px 12px rgba(16, 185, 129, 0.3);
}

.timeline-dot.current {
  background: linear-gradient(135deg, #fbbf24 0%, #f59e0b 100%);
  border-color: #f59e0b;
  color: #fff;
  box-shadow: 0 0 0 4px rgba(245, 158, 11, 0.2), 0 3px 12px rgba(245, 158, 11, 0.3);
  animation: pulseIndicator 2s infinite;
  font-size: 22px;
}

.timeline-dot.rejected {
  background: linear-gradient(135deg, #ef4444 0%, #dc2626 100%);
  border-color: #ef4444;
  color: #fff;
  box-shadow: 0 3px 12px rgba(239, 68, 68, 0.3);
}

.timeline-dot.pending {
  background: #f8f8f8;
  border-color: #e8e8e8;
  color: #ccc;
  box-shadow: 0 2px 6px rgba(0, 0, 0, 0.05);
}

@keyframes pulseIndicator {
  0%,
  100% {
    box-shadow: 0 0 0 4px rgba(245, 158, 11, 0.2), 0 3px 12px rgba(245, 158, 11, 0.3);
    transform: scale(1);
  }

  50% {
    box-shadow: 0 0 0 8px rgba(245, 158, 11, 0.1), 0 3px 12px rgba(245, 158, 11, 0.3);
    transform: scale(1.02);
  }
}

/* Timeline Label */
.timeline-label {
  font-size: 13px;
  color: #555;
  font-weight: 500;
  text-align: center;
  word-break: break-word;
  line-height: 1.4;
  letter-spacing: 0.2px;
  max-width: 90px;
}

/* Timeline Connector */
.timeline-connector {
  position: absolute;
  top: 19px;
  left: calc(50% + 22px);
  right: calc(-50% + 22px);
  height: 3px;
  background: #e8e8e8;
  z-index: 1;
  transition: all 0.3s ease;
}

.timeline-connector.passed {
  background: linear-gradient(to right, #10b981 0%, #10b981 100%);
  box-shadow: 0 1px 4px rgba(16, 185, 129, 0.2);
}

.timeline-connector.rejected {
  background: linear-gradient(to right, #10b981 0%, #ef4444 100%);
}

.timeline-item:last-child .timeline-connector {
  display: none;
}

/* Timeline Info Box */
.timeline-info {
  margin-top: 20px;
}

.info-box {
  background: #f0f9ff;
  border: 1px solid #bae6fd;
  border-radius: 12px;
  padding: 16px 20px;
  display: flex;
  align-items: flex-start;
  gap: 12px;
}

.info-box svg {
  color: #0369a1;
  flex-shrink: 0;
  margin-top: 2px;
}

.info-box p {
  font-size: 14px;
  color: #0c4a6e;
  margin: 0;
  line-height: 1.6;
  font-weight: 500;
}

/* ========== ACTION BUTTONS ========== */
.action-buttons {
  display: flex;
  justify-content: flex-end;
  gap: 15px;
  margin-top: 30px;
  padding-top: 25px;
  border-top: 1px solid #eee;
}

.btn-secondary,
.btn-primary {
  padding: 13px 32px;
  border: none;
  border-radius: 12px;
  font-size: 16px;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.2s;
  display: inline-flex;
  align-items: center;
  gap: 8px;
}

.btn-secondary {
  background: #f0f0f0;
  color: #666;
}

.btn-secondary:hover {
  background: #e0e0e0;
  color: #333;
}

.btn-primary {
  background: linear-gradient(135deg, #037266 0%, #026b5b 100%);
  color: #fff;
  box-shadow: 0 4px 12px rgba(3, 114, 102, 0.2);
}

.btn-primary:hover:not(:disabled) {
  background: linear-gradient(135deg, #026b5b 0%, #025a4d 100%);
  transform: translateY(-2px);
  box-shadow: 0 6px 16px rgba(3, 114, 102, 0.3);
}

.btn-primary:disabled {
  background: #ccc;
  cursor: not-allowed;
  box-shadow: none;
  opacity: 0.6;
}

.btn-confirm {
  background: linear-gradient(135deg, #10b981 0%, #059669 100%);
}

.btn-confirm:hover {
  background: linear-gradient(135deg, #059669 0%, #047857 100%);
}

/* ========== STEP CONTENT ========== */
.step-content {
  animation: fadeIn 0.3s ease;
}

@keyframes fadeIn {
  from {
    opacity: 0;
    transform: translateY(10px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

/* ========== RESPONSIVE ========== */
@media (max-width: 768px) {
  .application-page {
    padding: 15px;
  }

  .job-info-card {
    flex-direction: column;
    text-align: center;
    padding: 20px;
  }

  .company-logo {
    width: 70px;
    height: 70px;
  }

  .progress-steps {
    padding: 0;
  }

  .step-label {
    font-size: 12px;
  }

  .step-number {
    width: 40px;
    height: 40px;
    font-size: 16px;
  }

  .step-line {
    margin: 0 10px;
  }

  .section-card,
  .application-summary-card {
    padding: 20px;
  }

  .review-item {
    flex-direction: column;
    align-items: flex-start;
    gap: 5px;
  }

  .review-label {
    min-width: auto;
  }

  .review-value {
    text-align: left;
  }

  .action-buttons {
    flex-direction: column-reverse;
  }

  .btn-secondary,
  .btn-primary {
    width: 100%;
    justify-content: center;
  }

  .confirmation-notice {
    flex-direction: column;
    text-align: center;
  }

  .job-position-display {
    flex-direction: column;
    text-align: center;
  }

  .summary-icon {
    width: 70px;
    height: 70px;
  }

  .summary-header h2 {
    font-size: 22px;
  }

  .status-timeline-card {
    padding: 20px;
  }

  .timeline-header {
    flex-direction: column;
    gap: 15px;
    align-items: flex-start;
  }

  .btn-track-status {
    width: 100%;
  }

  .timeline-container {
    padding: 0 5px;
  }

  .timeline-dot {
    width: 45px;
    height: 45px;
  }

  .timeline-label {
    font-size: 13px;
    max-width: 85px;
  }

  .timeline-connector {
    top: 22px;
  }
}

@media (max-width: 480px) {
  .page-header h1 {
    font-size: 24px;
  }

  .job-info-content h2 {
    font-size: 18px;
  }

  .progress-steps {
    flex-wrap: wrap;
    gap: 20px;
  }

  .step-line {
    display: none;
  }

  .step {
    flex: 1 1 calc(33.333% - 15px);
    min-width: 80px;
  }

  .application-summary-card {
    padding: 25px 20px;
  }

  .summary-header {
    margin-bottom: 30px;
    padding-bottom: 25px;
  }

  .summary-icon {
    width: 60px;
    height: 60px;
  }

  .summary-icon svg {
    width: 36px;
    height: 36px;
  }

  .summary-header h2 {
    font-size: 20px;
  }

  .summary-section {
    padding: 20px;
  }
}
</style>