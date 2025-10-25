<!-- frontend/src/components/ApplicationSidePanel.vue -->
<template>
  <div v-if="show" class="side-panel-overlay" @click.self="$emit('close')">
    <div class="side-panel">
      <button class="side-close" @click="$emit('close')">&times;</button>
      <div class="side-panel-inner">
        <div class="panel-header">
          <!-- <h2>รายละเอียดใบสมัครงาน</h2>
          <p class="panel-subtitle">ตรวจสอบข้อมูลการสมัครงานของคุณได้ที่นี่</p> -->
        </div>
        
        <!-- เรียกใช้ ApplicationSummaryCard Component -->
        <ApplicationSummaryCard
          :job-data="jobData"
          :application-data="applicationData"
          :application-date="applicationDate"
          header-title="ใบสมัครงาน"
          header-subtitle="รายละเอียดการสมัครงานของคุณ"
        />
      </div>
    </div>
  </div>
</template>

<script setup>
import { defineProps, defineEmits } from 'vue';
import ApplicationSummaryCard from './ApplicationSummaryCard.vue';

const props = defineProps({
  show: {
    type: Boolean,
    default: false
  },
  jobData: {
    type: Object,
    required: true
  },
  applicationData: {
    type: Object,
    required: true
  },
  applicationDate: {
    type: String,
    required: true
  }
});

const emit = defineEmits(['close']);
</script>

<style scoped>
@import url("https://fonts.googleapis.com/css2?family=Kanit:wght@400;700&display=swap");

* {
  box-sizing: border-box;
  font-family: "Kanit", "Prompt", "Sarabun", "Noto Sans Thai", sans-serif;
}

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
  z-index: 10;
}

.side-close:hover {
  background: #ffc800;
  transform: scale(1.05);
}

.side-panel-inner {
  padding: 20px 0;
  display: flex;
  flex-direction: column;
  gap: 20px;
}

.panel-header {
  text-align: center;
  margin-bottom: 10px;
}

.panel-header h2 {
  font-size: 24px;
  font-weight: 700;
  color: #222;
  margin: 0 0 8px 0;
}

.panel-subtitle {
  font-size: 14px;
  color: #666;
  margin: 0;
}

@media (max-width: 768px) {
  .side-panel {
    width: 100vw;
    max-width: 100%;
    border-radius: 0;
    padding: 30px 20px;
  }

  .side-panel-inner {
    padding: 10px 0;
  }
}

@media (max-width: 480px) {
  .side-panel {
    padding: 20px 15px;
  }

  .side-close {
    top: 15px;
    right: 15px;
    width: 30px;
    height: 30px;
    font-size: 18px;
  }
}
</style>