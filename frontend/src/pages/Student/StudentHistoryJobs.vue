<!-- frontend/src/pages/Student/StudentHistoryJobs.vue -->
<template>
  <DashboardLayout role="student">
    <section class="history-section container">
      <!-- ========== FILTERS & SEARCH ========== -->
      <div class="history-controls">
        <CustomDropdown 
          label="บริษัท" 
          v-model="selectedCompany" 
          :options="companies"
          width="280px"
        />
        
        <SearchBar 
          v-model="searchText" 
          placeholder="ค้นหาชื่อตำแหน่ง หรือบริษัท"
          :is-ascending="isAscending"
          @toggle-sort="toggleSort"
        />
      </div>

      <!-- ========== HISTORY CARDS ========== -->
      <div class="history-list">
        <div v-for="record in filteredRecords" :key="record.id" class="history-card">
          <div class="card-header">
            <div class="header-info">
              <h3 class="job-title">{{ record.jobTitle }}</h3>
              <p class="company-name">{{ record.company }}</p>
              <p class="application-date">สมัครเมื่อ {{ formatDate(record.applicationDate) }}</p>
            </div>
            <div class="status-badge" :class="getStatusClass(record.stages)">
              {{ getCurrentStage(record.stages) }}
            </div>
          </div>

          <div class="timeline-container">
            <div class="timeline-wrapper">
              <div v-for="(stage, index) in record.stages" :key="index" class="timeline-item">
                <div class="timeline-dot" :class="stage.status">
                  <svg v-if="stage.status === 'passed'" xmlns="http://www.w3.org/2000/svg" width="18" height="18"
                    viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="3" stroke-linecap="round"
                    stroke-linejoin="round">
                    <polyline points="20 6 9 17 4 12"></polyline>
                  </svg>
                  <span v-else-if="stage.status === 'current'">?</span>
                  <svg v-else-if="stage.status === 'rejected'" xmlns="http://www.w3.org/2000/svg" width="18" height="18"
                    viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round"
                    stroke-linejoin="round">
                    <line x1="18" y1="6" x2="6" y2="18"></line>
                    <line x1="6" y1="6" x2="18" y2="18"></line>
                  </svg>
                </div>
                <div class="timeline-label">{{ stage.name }}</div>
                <div v-if="index < record.stages.length - 1" class="timeline-connector"
                  :class="getConnectorClass(record.stages, index)">
                </div>
              </div>
            </div>
          </div>

          <div class="card-actions">
            <button class="btn-detail" @click="openPanel(record)">ดูประกาศงาน</button>
            <button class="btn-application">ดูใบสมัครงาน</button>
          </div>
        </div>
      </div>
    </section>

    <!-- ========== JOB DETAIL SIDE PANEL ========== -->
    <div v-if="selectedJob" class="side-panel-overlay" @click.self="closePanel">
      <div class="side-panel">
        <button class="side-close" @click="closePanel">&times;</button>
        <div class="side-panel-inner">
          <div class="side-header-top">
            <img class="side-logo" :src="selectedJob.logo" alt="logo" />
            <div class="side-title-block">
              <div class="side-title">{{ selectedJob.title }}</div>
              <div class="side-company-row">
                <span class="side-company-name">{{ selectedJob.company }}</span>
                <span class="side-company-sep"></span>
                <span class="side-company-all-jobs">ดูงานทั้งหมด</span>
              </div>
              <span class="side-department-label">
                {{ selectedJob.department }}
              </span>
            </div>
          </div>
          <div class="side-meta-row">
            <span class="side-meta-item">
              <svg width="22" height="19" fill="none" stroke="#9CA3AF" stroke-width="2" stroke-linecap="round"
                stroke-linejoin="round" viewBox="0 0 24 24">
                <rect x="3" y="4" width="18" height="18" rx="2" ry="2" />
                <line x1="16" y1="2" x2="16" y2="6" />
                <line x1="8" y1="2" x2="8" y2="6" />
                <line x1="3" y1="10" x2="21" y2="10" />
              </svg>
              ประกาศเมื่อ {{ selectedJob.postDate }}
            </span>
          </div>
          <div class="side-meta-row">
            <span class="side-meta-item">
              <svg width="22" height="19" fill="none" stroke="#9CA3AF" stroke-width="2" stroke-linecap="round"
                stroke-linejoin="round" viewBox="0 0 24 24">
                <circle cx="12" cy="12" r="10"></circle>
                <polyline points="12 6 12 12 16 14"></polyline>
              </svg>
              ปิดรับสมัคร {{ selectedJob.closeDate }} (ระยะเวลา
              {{ selectedJob.applyDuration }})
            </span>
          </div>
          <div class="side-meta-row">
            <span class="side-meta-item">
              <svg width="22" height="22" fill="none" stroke="#9CA3AF" stroke-width="2" stroke-linecap="round"
                stroke-linejoin="round" viewBox="0 0 24 24">
                <path d="M22 10L12 4 2 10l10 6 10-6z"></path>
                <path d="M6 12v5c0 1.1 2.7 2 6 2s6-.9 6-2v-5"></path>
              </svg>
              ปีการศึกษา {{ selectedJob.academicYear }} ภาคที่
              {{ selectedJob.semester }}
            </span>
          </div>

          <!-- Images Preview -->
          <div v-if="selectedJob.images && selectedJob.images.length" class="side-images-preview">
            <div class="side-image-layout">
              <div class="side-image-large">
                <img :src="selectedJob.images[selectedImageIndex]" alt="job-image-large" />
              </div>
              <div class="side-image-thumbs-wrapper">
                <div class="side-image-thumbs">
                  <div v-for="(img, i) in selectedJob.images" :key="i" class="side-image-thumb"
                    :class="{ active: selectedImageIndex === i }" @click.stop="selectedImageIndex = i">
                    <img :src="img" alt="job-thumb" />
                  </div>
                </div>
              </div>
            </div>
          </div>
          <div class="side-detail-content">
            <div class="side-detail-title">รายละเอียดงาน</div>
            <div class="side-description">{{ selectedJob.description }}</div>
          </div>
        </div>
      </div>
    </div>
  </DashboardLayout>
</template>

<script setup>
import { ref, computed } from "vue";
import DashboardLayout from '../../components/DashboardLayout.vue';
import CustomDropdown from '../../components/CustomDropdown.vue';
import SearchBar from '../../components/SearchBar.vue';
import jobsList from '../../data/jobsData.js';

const selectedCompany = ref("ทั้งหมด");
const searchText = ref("");
const isAscending = ref(true);
const selectedJob = ref(null);
const selectedImageIndex = ref(0);

const monthsThai = [
  "มกราคม", "กุมภาพันธ์", "มีนาคม", "เมษายน", "พฤษภาคม", "มิถุนายน",
  "กรกฎาคม", "สิงหาคม", "กันยายน", "ตุลาคม", "พฤศจิกายน", "ธันวาคม"
];

// ใช้ข้อมูลจาก jobsData.js โดยสร้าง mapping ระหว่าง jobId กับสถานะการสมัคร
const applicationRecords = ref([
  {
    id: 1,
    jobId: 1, // อ้างอิงไปยัง job id ใน jobsData.js
    applicationDate: "01/01/68",
    stages: [
      { name: "กดสมัคร", status: "passed" },
      { name: "ตรวจสอบเอกสาร", status: "current" },
      { name: "สัมภาษณ์", status: "pending" },
      { name: "รอพิจารณา", status: "pending" },
      { name: "ประกาศผล", status: "pending" }
    ]
  },
  {
    id: 2,
    jobId: 2,
    applicationDate: "05/02/68",
    stages: [
      { name: "กดสมัคร", status: "passed" },
      { name: "ตรวจสอบเอกสาร", status: "passed" },
      { name: "สัมภาษณ์", status: "current" },
      { name: "รอพิจารณา", status: "pending" },
      { name: "ประกาศผล", status: "pending" }
    ]
  },
  {
    id: 3,
    jobId: 3,
    applicationDate: "10/03/68",
    stages: [
      { name: "กดสมัคร", status: "passed" },
      { name: "ตรวจสอบเอกสาร", status: "passed" },
      { name: "สัมภาษณ์", status: "passed" },
      { name: "รอพิจารณา", status: "current" },
      { name: "ประกาศผล", status: "pending" }
    ]
  },
  {
    id: 4,
    jobId: 4,
    applicationDate: "20/04/68",
    stages: [
      { name: "กดสมัคร", status: "passed" },
      { name: "ตรวจสอบเอกสาร", status: "rejected" },
      { name: "สัมภาษณ์", status: "pending" },
      { name: "รอพิจารณา", status: "pending" },
      { name: "ประกาศผล", status: "pending" }
    ]
  },
  {
    id: 5,
    jobId: 5,
    applicationDate: "01/05/68",
    stages: [
      { name: "กดสมัคร", status: "passed" },
      { name: "ตรวจสอบเอกสาร", status: "passed" },
      { name: "สัมภาษณ์", status: "passed" },
      { name: "รอพิจารณา", status: "passed" },
      { name: "ประกาศผล", status: "passed" }
    ]
  },
  {
    id: 6,
    jobId: 9,
    applicationDate: "15/09/68",
    stages: [
      { name: "กดสมัคร", status: "passed" },
      { name: "ตรวจสอบเอกสาร", status: "passed" },
      { name: "สัมภาษณ์", status: "current" },
      { name: "รอพิจารณา", status: "pending" },
      { name: "ประกาศผล", status: "pending" }
    ]
  }
]);

// รวมข้อมูลจาก applicationRecords กับ jobsList
const enrichedRecords = computed(() => {
  return applicationRecords.value.map(record => {
    const job = jobsList.find(j => j.id === record.jobId);
    return {
      ...record,
      jobTitle: job?.title || 'ไม่พบข้อมูล',
      company: job?.company || 'ไม่พบข้อมูล',
      department: job?.department || '',
      logo: job?.logo || '',
      postDate: job?.postDate || '',
      closeDate: job?.closeDate || '',
      applyDuration: job?.applyDuration || '',
      academicYear: job?.academicYear || '',
      semester: job?.semester || '',
      description: job?.description || '',
      images: job?.images || [],
      title: job?.title || ''
    };
  });
});

// สร้างรายการบริษัททั้งหมดจาก enrichedRecords
const companies = computed(() => {
  const uniqueCompanies = new Set(enrichedRecords.value.map(r => r.company));
  return ["ทั้งหมด", ...Array.from(uniqueCompanies)];
});

// ฟังก์ชันสำหรับ toggle การเรียงลำดับ
const toggleSort = () => {
  isAscending.value = !isAscending.value;
};

const formatDate = (dateStr) => {
  const parts = dateStr.split("/");
  const day = parts[0];
  const month = parseInt(parts[1]) - 1;
  const year = parseInt(parts[2]) + 2500;
  return `${day} ${monthsThai[month]} ${year}`;
};

const getCurrentStage = (stages) => {
  const currentStage = stages.find(s => s.status === "current");
  const rejectedStage = stages.find(s => s.status === "rejected");

  if (rejectedStage) {
    return `ไม่ผ่าน${rejectedStage.name}`;
  }

  if (currentStage) {
    return `${currentStage.name}`;
  }

  const passedStages = stages.filter(s => s.status === "passed");
  if (passedStages.length === stages.length) {
    return "ผ่านการคัดเลือก";
  }

  return "รอดำเนินการ";
};

const getStatusClass = (stages) => {
  const rejectedStage = stages.find(s => s.status === "rejected");
  if (rejectedStage) return "rejected";

  const currentStage = stages.find(s => s.status === "current");
  const passedStages = stages.filter(s => s.status === "passed");

  if (passedStages.length === stages.length) return "completed";

  if (currentStage) {
    const currentIndex = stages.findIndex(s => s.status === "current");
    if (currentIndex === 0 || currentIndex === 1) return "applying";
    if (currentIndex === 2) return "interview";
    if (currentIndex === 3) return "considering";
  }

  return "applying";
};

const getConnectorClass = (stages, index) => {
  const currentStage = stages[index];
  const nextStage = stages[index + 1];

  if (currentStage.status === "passed" && (nextStage.status === "passed" || nextStage.status === "current")) {
    return "passed";
  }

  if (currentStage.status === "passed" && nextStage.status === "rejected") {
    return "rejected";
  }

  return "pending";
};

const filteredRecords = computed(() => {
  // กรองตามบริษัทและคำค้นหา
  let filtered = enrichedRecords.value.filter(record => {
    if (selectedCompany.value !== "ทั้งหมด" && record.company !== selectedCompany.value) return false;
    if (searchText.value.trim()) {
      const txt = searchText.value.toLowerCase();
      return record.jobTitle.toLowerCase().includes(txt) || record.company.toLowerCase().includes(txt);
    }
    return true;
  });

  // เรียงลำดับตามวันที่สมัคร
  return isAscending.value ? filtered : [...filtered].reverse();
});

const openPanel = (record) => {
  selectedJob.value = record;
  selectedImageIndex.value = 0;
};

const closePanel = () => {
  selectedJob.value = null;
};
</script>

<style scoped>
@import url("https://fonts.googleapis.com/css2?family=Kanit:wght@400;500;600;700&display=swap");

.history-section {
  padding: 0;
}

* {
  box-sizing: border-box;
  font-family: "Kanit", "Prompt", "Sarabun", "Noto Sans Thai", sans-serif;
}

.history-controls {
  display: grid;
  grid-template-columns: 280px 1fr;
  gap: 16px;
  margin-bottom: 24px;
  align-items: end;
}

/* Responsive */
@media (max-width: 900px) {
  .history-controls {
    grid-template-columns: 1fr;
  }
}

.filter-row {
  display: flex;
  gap: 12px;
  align-items: flex-end;
  flex-wrap: wrap;
}

/* ========== HISTORY CARDS ========== */
.history-list {
  display: flex;
  flex-direction: column;
  gap: 18px;
  margin-top: 20px;
}

.history-card {
  background: #fff;
  border-radius: 16px;
  border: 1px solid #f0f0f0;
  padding: 26px 30px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.04);
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
}

.history-card:hover {
  box-shadow: 0 8px 24px rgba(3, 114, 102, 0.12);
  transform: translateY(-2px);
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: flex-start;
  gap: 20px;
  margin-bottom: 28px;
}

.header-info {
  flex: 1;
}

.job-title {
  font-size: 19px;
  font-weight: 700;
  color: #1a1a1a;
  margin: 0 0 8px 0;
  letter-spacing: -0.2px;
  line-height: 1.3;
}

.company-name {
  font-size: 14px;
  color: #666;
  font-weight: 500;
  margin: 0 0 6px 0;
  letter-spacing: 0.1px;
}

.application-date {
  font-size: 13px;
  color: #999;
  margin: 0;
  font-weight: 400;
}

/* ========== STATUS BADGES ========== */
.status-badge {
  padding: 8px 20px;
  border-radius: 20px;
  font-size: 13px;
  font-weight: 600;
  white-space: nowrap;
  transition: all 0.3s ease;
  flex-shrink: 0;
}

.status-badge.applying {
  background: #dbeafe;
  color: #1e40af;
}

.status-badge.interview {
  background: #e0f2fe;
  color: #0369a1;
}

.status-badge.considering {
  background: #fef3c7;
  color: #92400e;
}

.status-badge.completed {
  background: #d1fae5;
  color: #065f46;
}

.status-badge.rejected {
  background: #fee2e2;
  color: #991b1b;
}

/* ========== TIMELINE ========== */
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
  0%, 100% {
    box-shadow: 0 0 0 4px rgba(245, 158, 11, 0.2), 0 3px 12px rgba(245, 158, 11, 0.3);
    transform: scale(1);
  }
  50% {
    box-shadow: 0 0 0 8px rgba(245, 158, 11, 0.1), 0 3px 12px rgba(245, 158, 11, 0.3);
    transform: scale(1.02);
  }
}

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

/* ========== ACTION BUTTONS ========== */
.card-actions {
  display: flex;
  gap: 12px;
}

.btn-detail,
.btn-application {
  flex: 1;
  padding: 12px 22px;
  border: none;
  border-radius: 12px;
  font-size: 14px;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.25s ease;
  letter-spacing: 0.3px;
}

.btn-detail {
  background: linear-gradient(135deg, #f0f9f7 0%, #e8f5f2 100%);
  color: #037266;
  border: 2px solid #c5e4e0;
}

.btn-detail:hover {
  background: linear-gradient(135deg, #e0f2f0 0%, #d8eae7 100%);
  border-color: #037266;
  transform: translateY(-2px);
  box-shadow: 0 6px 16px rgba(3, 114, 102, 0.2);
}

.btn-detail:active {
  transform: translateY(0);
}

.btn-application {
  background: #fff;
  color: #666;
  border: 2px solid #e0e0e0;
}

.btn-application:hover {
  background: #f9f9f9;
  color: #333;
  border-color: #999;
  transform: translateY(-2px);
  box-shadow: 0 6px 16px rgba(0, 0, 0, 0.1);
}

.btn-application:active {
  transform: translateY(0);
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
  width: 697px;
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
  background: #ffd43b;
  color: #969696;
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
  background: #ffc800;
  transform: scale(1.05);
}

.side-panel-inner {
  padding: 46px 34px 38px 34px;
  display: flex;
  flex-direction: column;
  gap: 18px;
  min-height: 100vh;
}

.side-header-top {
  display: flex;
  align-items: flex-start;
  gap: 26px;
  margin-bottom: 4px;
}

.side-logo {
  width: 120px;
  height: 120px;
  border-radius: 50%;
  object-fit: cover;
  background: #f3f3f4;
  border: 2.2px solid #e1e1e1;
  flex-shrink: 0;
  box-shadow: 0 2px 13px rgba(8, 12, 22, 0.11);
}

.side-title-block {
  display: flex;
  flex-direction: column;
  gap: 7px;
  flex: 1;
  min-width: 0;
}

.side-title {
  font-size: 24px;
  font-weight: 700;
  color: #222;
  margin-bottom: 2px;
  margin-top: 0;
  word-break: break-word;
}

.side-company-row {
  display: flex;
  align-items: center;
  gap: 10px;
  font-size: 14px;
  color: #666;
  margin-bottom: 0;
}

.side-company-name {
  font-weight: 600;
  font-size: 15px;
  color: #037266;
  text-transform: uppercase;
  letter-spacing: 0.04em;
}

.side-company-sep {
  width: 5px;
  height: 5px;
  border-radius: 50%;
  background: #d7d7d7;
  display: inline-block;
}

.side-company-all-jobs {
  color: #888;
  font-size: 14px;
}

.side-department-label {
  display: inline-block;
  background-color: #cee5e2;
  color: #037266;
  border-radius: 7px;
  font-size: 13px;
  font-weight: 500;
  padding: 3px 14px;
  letter-spacing: 0.04em;
  margin-top: 7px;
  align-self: flex-start;
}

.side-meta-row {
  display: flex;
  align-items: center;
  gap: 0px;
  margin: 7px 0 0 0;
  font-size: 14px;
  color: #667;
  flex-wrap: wrap;
  flex-direction: row;
}

.side-meta-row+.side-meta-row {
  margin-top: 0;
}

.side-meta-item {
  display: flex;
  align-items: center;
  gap: 7px;
  font-size: 15px;
  color: #3b5454;
  font-weight: 500;
  margin-bottom: 0.5px;
}

.side-meta-item svg {
  flex-shrink: 0;
  margin-bottom: -2px;
}

.side-images-preview {
  margin: 25px 0 11px 0;
  display: flex;
  flex-direction: row;
  gap: 11px;
}

.side-image-layout {
  display: flex;
  flex-direction: row;
  gap: 16px;
}

.side-image-large {
  width: 360px;
  height: 420px;
  background: #f3f3f4;
  border-radius: 13px;
  border: 1.5px solid #e1e1e1;
  display: flex;
  align-items: center;
  justify-content: center;
  box-shadow: 0 2px 11px rgba(8, 12, 22, 0.09);
  flex-shrink: 0;
}

.side-image-large img {
  width: 95%;
  height: 95%;
  object-fit: cover;
  border-radius: 10px;
  background: #fff;
  box-shadow: 0 1px 8px rgba(0, 0, 0, 0.07);
}

.side-image-thumbs-wrapper {
  height: 420px;
  max-height: 420px;
  overflow-y: auto;
  overflow-x: hidden;
  flex-shrink: 0;
}

.side-image-thumbs-wrapper::-webkit-scrollbar {
  width: 6px;
}

.side-image-thumbs-wrapper::-webkit-scrollbar-track {
  background: #f1f1f1;
  border-radius: 10px;
}

.side-image-thumbs-wrapper::-webkit-scrollbar-thumb {
  background: #c1c1c1;
  border-radius: 10px;
}

.side-image-thumbs-wrapper::-webkit-scrollbar-thumb:hover {
  background: #a8a8a8;
}

.side-image-thumbs {
  display: flex;
  flex-direction: column;
  gap: 12px;
  align-items: flex-start;
  justify-content: flex-start;
  min-width: 96px;
}

.side-image-thumb {
  width: 96px;
  height: 96px;
  border-radius: 8px;
  overflow: hidden;
  border: 2px solid transparent;
  background: #f3f3f4;
  cursor: pointer;
  box-shadow: 0 1px 6px rgba(8, 12, 22, 0.09);
  transition: border-color 0.18s, box-shadow 0.18s, transform 0.15s;
  display: flex;
  align-items: center;
  justify-content: center;
  flex-shrink: 0;
}

.side-image-thumb img {
  width: 100%;
  height: 100%;
  object-fit: cover;
  border-radius: 6px;
}

.side-image-thumb.active {
  border-color: #037266;
  box-shadow: 0 3px 12px rgba(3, 114, 102, 0.18);
  transform: scale(1.03);
}

.side-image-thumb:hover {
  border-color: #037266;
  box-shadow: 0 3px 12px rgba(3, 114, 102, 0.14);
}

.side-detail-content {
  margin-top: 18px;
  display: flex;
  flex-direction: column;
  gap: 15px;
}

.side-detail-title {
  font-size: 18px;
  font-weight: 700;
  color: #222;
  margin-bottom: 3px;
}

.side-description {
  font-size: 15px;
  line-height: 1.7;
  color: #444;
  background: none;
  padding: 0;
  border-radius: 0;
  white-space: pre-line;
}

/* Responsive */
@media (max-width: 900px) {
  .history-card {
    padding: 20px 24px;
  }

  .card-header {
    flex-direction: column;
    gap: 14px;
  }

  .status-badge {
    align-self: flex-start;
  }

  .side-panel-inner {
    padding: 28px 20px 21px 20px;
  }

  .side-image-layout {
    flex-direction: column;
    gap: 16px;
    align-items: center;
  }

  .side-image-large {
    width: 100%;
    max-width: 400px;
    height: 300px;
  }

  .side-image-thumbs-wrapper {
    height: auto;
    max-height: none;
    overflow-y: visible;
    overflow-x: auto;
    width: 100%;
    max-width: 400px;
  }

  .side-image-thumbs {
    flex-direction: row;
    gap: 10px;
    width: max-content;
  }

  .side-image-thumb {
    width: 80px;
    height: 80px;
  }
}

@media (max-width: 650px) {
  .history-card {
    padding: 16px 18px;
  }

  .job-title {
    font-size: 16px;
  }

  .status-badge {
    padding: 8px 18px;
    font-size: 13px;
  }

  .timeline-container {
    padding: 0 5px;
  }

  .timeline-wrapper {
    gap: 3px;
  }

  .timeline-dot {
    width: 36px;
    height: 36px;
  }

  .timeline-dot svg {
    width: 16px;
    height: 16px;
  }

  .timeline-dot.current {
    font-size: 18px;
  }

  .timeline-label {
    font-size: 11px;
    max-width: 70px;
  }

  .timeline-connector {
    top: 17px;
    height: 2px;
    left: calc(50% + 20px);
    right: calc(-50% + 20px);
  }

  .card-actions {
    flex-direction: column;
    gap: 10px;
  }

  .btn-detail,
  .btn-application {
    width: 100%;
  }

  .side-panel {
    width: 100vw;
    min-width: 0;
    padding: 0;
    border-radius: 0;
  }

  .side-close {
    top: 12px;
    right: 10px;
    width: 32px;
    height: 32px;
    font-size: 21px;
  }

  .side-logo {
    width: 70px;
    height: 70px;
  }

  .side-panel-inner {
    padding: 14px 12px 10px 12px;
  }

  .side-image-layout {
    flex-direction: column;
    gap: 12px;
    align-items: center;
  }

  .side-image-large {
    width: 100%;
    max-width: 100%;
    height: 240px;
  }

  .side-image-thumbs-wrapper {
    height: auto;
    max-height: none;
    overflow-y: visible;
    overflow-x: auto;
    width: 100%;
    max-width: 100%;
  }

  .side-image-thumbs {
    flex-direction: row;
    gap: 8px;
    width: max-content;
  }

  .side-image-thumb {
    width: 65px;
    height: 65px;
  }
}
</style>