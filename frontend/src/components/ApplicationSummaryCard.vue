<!-- frontend/src/components/ApplicationSummaryCard.vue -->
<template>
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
      <h2>{{ headerTitle }}</h2>
      <p class="summary-subtitle">{{ headerSubtitle }}</p>
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
              <span class="document-name">{{ applicationData.resumeName }}</span>
            </div>
          </div>

          <!-- Cover Letter (if exists) -->
          <div v-if="applicationData.coverLetterName" class="document-item">
            <div class="document-icon cover-icon">
              <svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none"
                stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                <path d="M14 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8z"></path>
                <polyline points="14 2 14 8 20 8"></polyline>
              </svg>
            </div>
            <div class="document-info">
              <span class="document-type">จดหมายสมัครงาน (Cover Letter)</span>
              <span class="document-name">{{ applicationData.coverLetterName }}</span>
            </div>
          </div>

          <!-- Transcript (if exists) -->
          <div v-if="applicationData.transcriptName" class="document-item">
            <div class="document-icon transcript-icon">
              <svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none"
                stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                <path d="M14 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8z"></path>
                <polyline points="14 2 14 8 20 8"></polyline>
              </svg>
            </div>
            <div class="document-info">
              <span class="document-type">ใบแสดงผลการเรียน (Transcript)</span>
              <span class="document-name">{{ applicationData.transcriptName }}</span>
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
      <span>สมัครเมื่อวันที่ {{ applicationDate }}</span>
    </div>
  </div>
</template>

<script setup>
import { defineProps } from 'vue';

const props = defineProps({
  jobData: {
    type: Object,
    required: true,
    default: () => ({
      title: '',
      company: '',
      department: '',
      logo: ''
    })
  },
  applicationData: {
    type: Object,
    required: true,
    default: () => ({
      fullName: '',
      email: '',
      phone: '',
      resumeName: '',
      coverLetterName: null,
      transcriptName: null
    })
  },
  applicationDate: {
    type: String,
    required: true,
    default: ''
  },
  headerTitle: {
    type: String,
    default: 'ใบสมัครงานสำเร็จเรียบร้อยแล้ว'
  },
  headerSubtitle: {
    type: String,
    default: 'กรุณาตรวจสอบรายละเอียดใบสมัครของคุณ'
  }
});
</script>

<style scoped>
@import url("https://fonts.googleapis.com/css2?family=Kanit:wght@400;500;600;700&display=swap");

* {
  font-family: "Kanit", sans-serif;
  box-sizing: border-box;
}

.application-summary-card {
  background: #f9f9f9;
  border-radius: 16px;
  padding: 40px;
  box-shadow: 0 2px 12px rgba(3, 114, 102, 0.07);
  border: 1px solid #eef0ef;
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

.transcript-icon {
  background: linear-gradient(135deg, #3b82f6 0%, #2563eb 100%);
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

/* Responsive */
@media (max-width: 768px) {
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

  .job-position-display {
    flex-direction: column;
    text-align: center;
  }
}

@media (max-width: 480px) {
  .application-summary-card {
    padding: 20px 15px;
  }

  .summary-header h2 {
    font-size: 18px;
  }
}
</style>