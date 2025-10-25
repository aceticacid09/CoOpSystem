<!-- frontend/src/pages/Student/StudentDashboard.vue -->
<template>
  <DashboardLayout role="student">
    <!-- ====================== Summary Cards ====================== -->
    <div class="summary-cards">
      <div class="stat-card">
        <div class="stat-icon job-icon">
          <svg xmlns="http://www.w3.org/2000/svg" fill="none" stroke="currentColor" stroke-width="2"
            viewBox="0 0 24 24">
            <rect x="2" y="7" width="20" height="14" rx="2" ry="2" />
            <path d="M16 21V5a2 2 0 0 0-2-2h-4a2 2 0 0 0-2 2v16" />
          </svg>
        </div>
        <div class="stat-content">
          <p class="stat-label">จำนวนงานที่สมัคร</p>
          <h2 class="stat-value">{{ summary.totalJobs }}</h2>
        </div>
      </div>

      <div class="stat-card">
        <div class="stat-icon selected-icon">
          <svg xmlns="http://www.w3.org/2000/svg" fill="none" stroke="currentColor" stroke-width="2"
            viewBox="0 0 24 24">
            <path d="M22 11.08V12a10 10 0 1 1-5.93-9.14" />
            <polyline points="22 4 12 14.01 9 11.01" />
          </svg>
        </div>
        <div class="stat-content">
          <p class="stat-label">ผ่านการคัดเลือก</p>
          <h2 class="stat-value">{{ summary.selected }}</h2>
        </div>
      </div>

      <div class="stat-card">
        <div class="stat-icon pending-icon">
          <svg xmlns="http://www.w3.org/2000/svg" fill="none" stroke="currentColor" stroke-width="2"
            viewBox="0 0 24 24">
            <circle cx="12" cy="12" r="10" />
            <polyline points="12 6 12 12 16 14" />
          </svg>
        </div>
        <div class="stat-content">
          <p class="stat-label">รอผลรอบคัดเลือก</p>
          <h2 class="stat-value">{{ summary.pending }}</h2>
        </div>
      </div>
    </div>

    <!-- ====================== Job Applications ====================== -->
    <div class="section-container">
      <div class="section-header">
        <h3 class="section-title">ประวัติการสมัครงาน</h3>
        <router-link to="/student/jobs/history" class="view-all-link">ดูทั้งหมด</router-link>
      </div>

      <div class="job-list scrollable-list">
        <div v-for="(app, index) in applications" :key="index" class="job-item">
          <div class="job-color-bar" :class="barClass(app.status)"></div>
          <div class="job-info">
            <h4 class="job-title">{{ app.title }}</h4>
            <p class="job-deadline">สมัครเมื่อ: {{ app.date }}</p>
          </div>
          <span class="job-status" :class="statusClass(app.status)">
            {{ statusText(app.status) }}
          </span>
        </div>
      </div>
    </div>

    <!-- ====================== Available Jobs ====================== -->
    <div class="section-container">
      <div class="section-header">
        <h3 class="section-title">รายการโปรด</h3>
        <router-link to="/student/favorites" class="view-all-link">ดูทั้งหมด</router-link>
      </div>

      <div class="job-list scrollable-list-jobs">
        <FavoriteSections :saved-news-ids="savedNewsIds" :saved-jobs-ids="savedJobsIds" :show-bookmark="true"
        :show-actions="true" @remove-news-bookmark="handleRemoveNewsBookmark"
        @remove-job-bookmark="handleRemoveJobBookmark" />
      </div>
    </div>
  </DashboardLayout>
</template>

<script setup>
import { ref } from 'vue';
import DashboardLayout from '../../components/DashboardLayout.vue';
import FavoriteSections from '../../components/FavoriteSections.vue';

/* ====================== Summary Data ====================== */
const summary = ref({
  totalJobs: 6,
  selected: 2,
  pending: 3
});

/* ====================== Applications Data ====================== */
const applications = ref([
  { title: 'Web Developer - ABC Ltd.', date: '15 มีนาคม 2026', status: 'accepted' },
  { title: 'Software Developer - บริษัท ABC จำกัด', date: '15 มีนาคม 2026', status: 'pending' },
  { title: 'Data Analyst - บริษัท XYZ จำกัด', date: '15 มีนาคม 2026', status: 'rejected' },
  { title: 'Mobile Developer - บริษัท Mobi จำกัด', date: '18 มีนาคม 2026', status: 'pending' },
  { title: 'UX/UI Designer - บริษัท Creative จำกัด', date: '20 มีนาคม 2026', status: 'accepted' },
  { title: 'QA Tester - บริษัท Quality จำกัด', date: '22 มีนาคม 2026', status: 'rejected' }
]);

/* ====================== Helper Functions ====================== */
const barClass = (status) => {
  switch (status) {
    case 'accepted': return 'accepted-bar';
    case 'pending': return 'pending-bar';
    case 'rejected': return 'rejected-bar';
    default: return '';
  }
};

const statusClass = (status) => {
  return status;
};

const statusText = (status) => {
  switch (status) {
    case 'accepted': return 'ผ่านการคัดเลือก';
    case 'pending': return 'รอผลรอบคัดเลือก';
    case 'rejected': return 'ไม่ผ่านการคัดเลือก';
    default: return '';
  }
};

const savedNewsIds = ref([1, 2, 4]);
const savedJobsIds = ref([1, 3, 5]);

// Handle remove bookmark
const handleRemoveNewsBookmark = (newsId) => {
  const index = savedNewsIds.value.indexOf(newsId);
  if (index > -1) {
    savedNewsIds.value.splice(index, 1);
  }
  console.log('Removed news bookmark:', newsId);
};

const handleRemoveJobBookmark = (jobId) => {
  const index = savedJobsIds.value.indexOf(jobId);
  if (index > -1) {
    savedJobsIds.value.splice(index, 1);
  }
  console.log('Removed job bookmark:', jobId);
};
</script>

<style scoped>
/* Summary Cards */
.summary-cards {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(280px, 1fr));
  gap: 20px;
  margin-bottom: 32px;
}

.stat-card {
  background: linear-gradient(135deg, #ffffff 0%, #f8f9fa 100%);
  border-radius: 16px;
  padding: 24px;
  display: flex;
  align-items: center;
  gap: 20px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.04);
  border: 1px solid #f0f0f0;
  transition: all 0.3s ease;
}

.stat-card:hover {
  transform: translateY(-4px);
  box-shadow: 0 8px 20px rgba(3, 114, 102, 0.1);
}

.stat-icon {
  width: 64px;
  height: 64px;
  border-radius: 14px;
  display: flex;
  align-items: center;
  justify-content: center;
  flex-shrink: 0;
}

.stat-icon svg {
  width: 32px;
  height: 32px;
}

.job-icon {
  background: linear-gradient(135deg, #037266 0%, #025952 100%);
  color: white;
}

.selected-icon {
  background: linear-gradient(135deg, #10b981 0%, #059669 100%);
  color: white;
}

.pending-icon {
  background: linear-gradient(135deg, #f59e0b 0%, #d97706 100%);
  color: white;
}

.stat-content {
  flex: 1;
}

.stat-label {
  font-size: 14px;
  color: #666;
  margin: 0 0 6px 0;
  font-weight: 500;
}

.stat-value {
  font-size: 36px;
  font-weight: 700;
  color: #037266;
  margin: 0;
  line-height: 1;
}

/* Section Container */
.section-container {
  background: white;
  border-radius: 16px;
  padding: 28px;
  margin-bottom: 28px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.04);
  border: 1px solid #f0f0f0;
}

.section-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 24px;
  padding-bottom: 16px;
  border-bottom: 2px solid #f5f5f5;
}

.section-title {
  font-size: 20px;
  font-weight: 600;
  color: #333;
  margin: 0;
}

.view-all-link {
  color: #037266;
  text-decoration: none;
  font-weight: 500;
  font-size: 14px;
  padding: 8px 16px;
  border-radius: 8px;
  transition: all 0.2s ease;
}

.view-all-link:hover {
  background: #f0f8f7;
  color: #025952;
}

/* Job List */
.job-list {
  display: flex;
  flex-direction: column;
  gap: 16px;
}

.scrollable-list {
  max-height: 400px;
  overflow-y: auto;
  padding-right: 6px;
}

.scrollable-list-jobs {
  max-height: 500px;
  overflow-y: auto;
  padding-right: 6px;
}

.job-item {
  display: flex;
  align-items: center;
  gap: 20px;
  padding: 20px;
  background: #fafafa;
  border-radius: 12px;
  transition: all 0.3s ease;
  border: 1px solid #f0f0f0;
}

.job-item:hover {
  background: #f5f9f8;
  transform: translateX(4px);
  box-shadow: 0 4px 12px rgba(3, 114, 102, 0.08);
}

.job-color-bar {
  width: 6px;
  height: 60px;
  border-radius: 3px;
  flex-shrink: 0;
}

.accepted-bar {
  background: linear-gradient(180deg, #10b981 0%, #059669 100%);
}

.pending-bar {
  background: linear-gradient(180deg, #f59e0b 0%, #d97706 100%);
}

.rejected-bar {
  background: linear-gradient(180deg, #ef4444 0%, #dc2626 100%);
}

.job-info {
  flex: 1;
}

.job-title {
  font-size: 16px;
  font-weight: 600;
  color: #333;
  margin: 0 0 6px 0;
}

.job-deadline {
  font-size: 13px;
  color: #888;
  margin: 0;
}

.job-status {
  padding: 8px 20px;
  border-radius: 20px;
  font-size: 13px;
  font-weight: 600;
  white-space: nowrap;
}

.job-status.accepted {
  background: #d1fae5;
  color: #065f46;
}

.job-status.pending {
  background: #fef3c7;
  color: #92400e;
}

.job-status.rejected {
  background: #fee2e2;
  color: #991b1b;
}

/* Job Cards */
.job-cards {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(300px, 1fr));
  gap: 20px;
}

.job-card {
  background: white;
  border: 2px solid #f0f0f0;
  border-radius: 14px;
  padding: 24px;
  transition: all 0.3s ease;
  display: flex;
  flex-direction: column;
  justify-content: space-between;
  min-height: 200px;
}

.job-card:hover {
  border-color: #037266;
  transform: translateY(-4px);
  box-shadow: 0 8px 24px rgba(3, 114, 102, 0.12);
}

.job-card-header {
  margin-bottom: 20px;
}

.job-card-title {
  font-size: 18px;
  font-weight: 600;
  color: #333;
  margin: 0 0 8px 0;
}

.company-name {
  font-size: 14px;
  color: #666;
  margin: 0;
  line-height: 1.5;
}

.job-card-footer {
  display: flex;
  justify-content: space-between;
  align-items: flex-end;
  gap: 16px;
}

.job-card-deadline {
  font-size: 13px;
  color: #888;
  margin: 0;
  line-height: 1.6;
}

.days-left {
  color: #037266;
  font-weight: 500;
}

.apply-btn {
  background: linear-gradient(135deg, #037266 0%, #025952 100%);
  color: white;
  border: none;
  padding: 10px 24px;
  border-radius: 10px;
  font-weight: 600;
  font-size: 14px;
  cursor: pointer;
  transition: all 0.3s ease;
  white-space: nowrap;
  box-shadow: 0 2px 8px rgba(3, 114, 102, 0.2);
  font-family: "Kanit", "Prompt", "Sarabun", "Noto Sans Thai", sans-serif;
}

.apply-btn:hover {
  transform: translateY(-2px);
  box-shadow: 0 6px 16px rgba(3, 114, 102, 0.3);
  background: linear-gradient(135deg, #025952 0%, #014842 100%);
}

.apply-btn:active {
  transform: translateY(0);
}

/* Responsive Design */
@media (max-width: 768px) {
  .summary-cards {
    grid-template-columns: 1fr;
  }

  .section-container {
    padding: 20px;
  }

  .job-item {
    flex-direction: column;
    align-items: flex-start;
    gap: 12px;
  }

  .job-color-bar {
    width: 100%;
    height: 4px;
  }

  .job-status {
    align-self: flex-start;
  }

  .job-cards {
    grid-template-columns: 1fr;
  }

  .job-card-footer {
    flex-direction: column;
    align-items: stretch;
    gap: 12px;
  }

  .apply-btn {
    width: 100%;
  }
}

@media (max-width: 480px) {
  .stat-card {
    padding: 18px;
  }

  .stat-icon {
    width: 52px;
    height: 52px;
  }

  .stat-icon svg {
    width: 26px;
    height: 26px;
  }

  .stat-value {
    font-size: 28px;
  }

  .section-title {
    font-size: 18px;
  }
}
</style>