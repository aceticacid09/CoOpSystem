<!-- frontend/src/pages/Company/CompanyDashboard.vue -->
<template>
  <DashboardLayout role="company">
    <!-- ====================== Manage Jobs Section ====================== -->
    <div class="section-container">
      <div class="section-header">
        <h3 class="section-title">จัดการประกาศงาน</h3>
        <router-link to="/company/jobs/manage" class="view-all-link">ดูทั้งหมด</router-link>
      </div>

      <!-- <div class="preview-content"> --> <!-- ปลดคอมเม้นเพื่อซ่อน fitter -->
        <ManageJobsSections :isPreview="true" :previewLimit="5" />
      <!-- </div> -->
    </div>

    <!-- ====================== Job Applications Section ====================== -->
    <div class="section-container">
      <div class="section-header">
        <h3 class="section-title">ใบสมัครงานล่าสุด</h3>
        <router-link to="/company/jobs/applications" class="view-all-link">ดูทั้งหมด</router-link>
      </div>

      <!-- <div class="preview-content"> -->
        <JobApplicationsList :isPreview="true" :previewLimit="5" />
      <!-- </div> -->
    </div>
  </DashboardLayout>
</template>

<script setup>
import { ref } from 'vue';
import DashboardLayout from '../../components/DashboardLayout.vue';
import ManageJobsSections from '../../components/ManageJobsSections.vue';
import JobApplicationsList from '../../components/JobApplicationsList.vue';

</script>

<style scoped>
@import url("https://fonts.googleapis.com/css2?family=Kanit:wght@400;500;600;700&display=swap");

* {
  font-family: "Kanit", "Prompt", "Sarabun", "Noto Sans Thai", sans-serif;
  box-sizing: border-box;
}

/* ====================== Section Container ====================== */
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

/* ====================== Preview Content ====================== */
/* .preview-content { ซ่อน filters และ search ในโหมด preview
} */

.preview-content :deep(.jobs-filters),
.preview-content :deep(.search-sort) {
  display: none;
}

.preview-content :deep(.pagination-wrapper) {
  display: none;
}

.preview-content :deep(.section-container) {
  background: transparent;
  padding: 0;
  margin: 0;
  box-shadow: none;
  border: none;
}

.preview-content :deep(.table-wrapper) {
  border-radius: 12px;
}

/* ====================== Responsive Design ====================== */
@media (max-width: 768px) {
  .summary-cards {
    grid-template-columns: repeat(2, 1fr);
    gap: 16px;
  }

  .section-container {
    padding: 20px;
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

  .section-title {
    font-size: 18px;
  }
}
</style>