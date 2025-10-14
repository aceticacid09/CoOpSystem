<!-- frontend/src/pages/Company/CompanyDashboard.vue -->
<template>
  <DashboardLayout role="company">
    <!-- ====================== Summary Cards ====================== -->
    <div class="summary-cards">
      <div class="stat-card">
        <div class="stat-icon posted-icon">
          <svg xmlns="http://www.w3.org/2000/svg" fill="none" stroke="currentColor" stroke-width="2"
            viewBox="0 0 24 24">
            <path d="M11 4H4a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2v-7" />
            <path d="M18.5 2.5a2.121 2.121 0 0 1 3 3L12 15l-4 1 1-4 9.5-9.5z" />
          </svg>
        </div>
        <div class="stat-content">
          <p class="stat-label">ประกาศงานทั้งหมด</p>
          <h2 class="stat-value">{{ companyStats.totalPostings }}</h2>
        </div>
      </div>

      <div class="stat-card">
        <div class="stat-icon applications-icon">
          <svg xmlns="http://www.w3.org/2000/svg" fill="none" stroke="currentColor" stroke-width="2"
            viewBox="0 0 24 24">
            <path d="M20 21v-2a4 4 0 0 0-4-4H8a4 4 0 0 0-4 4v2" />
            <circle cx="12" cy="7" r="4" />
          </svg>
        </div>
        <div class="stat-content">
          <p class="stat-label">ใบสมัครงานทั้งหมด</p>
          <h2 class="stat-value">{{ companyStats.totalApplications }}</h2>
        </div>
      </div>

      <div class="stat-card">
        <div class="stat-icon reviewing-icon">
          <svg xmlns="http://www.w3.org/2000/svg" fill="none" stroke="currentColor" stroke-width="2"
            viewBox="0 0 24 24">
            <circle cx="12" cy="12" r="10" />
            <polyline points="12 6 12 12 16 14" />
          </svg>
        </div>
        <div class="stat-content">
          <p class="stat-label">ผู้สมัครที่รอพิจารณา</p>
          <h2 class="stat-value">{{ companyStats.underReview }}</h2>
        </div>
      </div>

      <div class="stat-card">
        <div class="stat-icon approved-icon">
          <svg xmlns="http://www.w3.org/2000/svg" fill="none" stroke="currentColor" stroke-width="2"
            viewBox="0 0 24 24">
            <path d="M22 11.08V12a10 10 0 1 1-5.93-9.14" />
            <polyline points="22 4 12 14.01 9 11.01" />
          </svg>
        </div>
        <div class="stat-content">
          <p class="stat-label">ผู้สมัครที่ผ่านการคัดเลือก</p>
          <h2 class="stat-value">{{ companyStats.approved }}</h2>
        </div>
      </div>
    </div>
  </DashboardLayout>
</template>

<script setup>
import { ref } from 'vue';
import DashboardLayout from '../../components/DashboardLayout.vue';

const companyStats = ref({
  totalPostings: 8,
  totalApplications: 24,
  underReview: 12,
  approved: 6
});
</script>

<style scoped>
@import url("https://fonts.googleapis.com/css2?family=Kanit:wght@400;500;600;700&display=swap");

* {
  font-family: "Kanit", "Prompt", "Sarabun", "Noto Sans Thai", sans-serif;
  box-sizing: border-box;
}

/* ====================== Summary Cards ====================== */
.summary-cards {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
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

.posted-icon {
  background: linear-gradient(135deg, #037266 0%, #025952 100%);
  color: white;
}

.applications-icon {
  background: linear-gradient(135deg, #0369a1 0%, #075985 100%);
  color: white;
}

.reviewing-icon {
  background: linear-gradient(135deg, #f59e0b 0%, #d97706 100%);
  color: white;
}

.approved-icon {
  background: linear-gradient(135deg, #10b981 0%, #059669 100%);
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

/* ====================== Responsive Design ====================== */
@media (max-width: 768px) {
  .summary-cards {
    grid-template-columns: repeat(2, 1fr);
    gap: 16px;
  }
}

@media (max-width: 480px) {
  .summary-cards {
    grid-template-columns: 1fr;
    gap: 12px;
  }

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
}
</style>
