<!-- frontend/src/components/JobSidePanel.vue -->
<template>
  <div v-if="job" class="side-panel-overlay" @click.self="$emit('close')">
    <div class="side-panel">
      <button class="side-close" @click="$emit('close')">&times;</button>
      <div class="side-panel-inner">
        <div class="side-header-top">
          <img class="side-logo" :src="job.logo" alt="logo" />
          <div class="side-title-block">
            <div class="side-title">{{ job.title }}</div>
            <div class="side-company-row">
              <span class="side-company-name">{{ job.company }}</span>
              <span class="side-company-sep"></span>
              <span class="side-company-type">{{ companyTypeText }}</span>
            </div>
            <span class="side-department-label">{{ job.department }}</span>
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
            ประกาศเมื่อ {{ job.postDate }}
          </span>
        </div>

        <div class="side-meta-row">
          <span class="side-meta-item">
            <svg width="22" height="19" fill="none" stroke="#9CA3AF" stroke-width="2" stroke-linecap="round"
              stroke-linejoin="round" viewBox="0 0 24 24">
              <circle cx="12" cy="12" r="10"></circle>
              <polyline points="12 6 12 12 16 14"></polyline>
            </svg>
            ปิดรับสมัคร {{ job.closeDate }} <span v-if="job.applyDuration">(ระยะเวลา {{ job.applyDuration }})</span>
          </span>
        </div>

        <div v-if="job.academicYear && job.semester" class="side-meta-row">
          <span class="side-meta-item">
            <svg width="22" height="22" fill="none" stroke="#9CA3AF" stroke-width="2" stroke-linecap="round"
              stroke-linejoin="round" viewBox="0 0 24 24">
              <path d="M22 10L12 4 2 10l10 6 10-6z"></path>
              <path d="M6 12v5c0 1.1 2.7 2 6 2s6-.9 6-2v-5"></path>
            </svg>
            ปีการศึกษา {{ job.academicYear }} ภาคที่ {{ job.semester }}
          </span>
        </div>

        <!-- Action Buttons -->
        <div v-if="showActions" class="side-action-buttons">
          <button class="btn-apply" :class="{ applied: isApplied }" @click="handleApply" :disabled="isApplied">
            <svg v-if="!isApplied" xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24"
              fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
              <path d="M16 21v-2a4 4 0 0 0-4-4H6a4 4 0 0 0-4 4v2"></path>
              <circle cx="9" cy="7" r="4"></circle>
              <line x1="19" y1="8" x2="19" y2="14"></line>
              <line x1="22" y1="11" x2="16" y2="11"></line>
            </svg>
            <svg v-else xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none"
              stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
              <path d="M22 11.08V12a10 10 0 1 1-5.93-9.14"></path>
              <polyline points="22 4 12 14.01 9 11.01"></polyline>
            </svg>
            <span>{{ isApplied ? 'สมัครงานแล้ว' : 'สมัครงาน' }}</span>
          </button>

          <button class="btn-save" :class="{ saved: isSaved }" @click="$emit('toggle-save', job.id)">
            <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24"
              :fill="isSaved ? 'currentColor' : 'none'" stroke="currentColor" stroke-width="2" stroke-linecap="round"
              stroke-linejoin="round">
              <path d="M19 21l-7-5-7 5V5a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2z"></path>
            </svg>
            <span>{{ isSaved ? 'บันทึกแล้ว' : 'บันทึก' }}</span>
          </button>
        </div>

        <!-- Images Preview -->
        <div v-if="job.images && job.images.length" class="side-images-preview">
          <div class="side-image-layout">
            <div class="side-image-large">
              <img :src="job.images[selectedImageIndex]" alt="job-image-large" />
            </div>
            <div class="side-image-thumbs-wrapper">
              <div class="side-image-thumbs">
                <div v-for="(img, i) in job.images" :key="i" class="side-image-thumb"
                  :class="{ active: selectedImageIndex === i }" @click.stop="selectedImageIndex = i">
                  <img :src="img" alt="job-thumb" />
                </div>
              </div>
            </div>
          </div>
        </div>

        <div class="side-detail-content">
          <div class="side-detail-title">รายละเอียดงาน</div>
          <div class="side-description">{{ job.description }}</div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, watch } from 'vue';

const props = defineProps({
  job: {
    type: Object,
    default: null
  },
  showActions: {
    type: Boolean,
    default: true
  },
  isApplied: {
    type: Boolean,
    default: false
  },
  isSaved: {
    type: Boolean,
    default: false
  },
  companyTypeText: {
    type: String,
    default: 'ดูรายละเอียดงาน'
  }
});

const emit = defineEmits(['close', 'apply', 'toggle-save']);

const selectedImageIndex = ref(0);

const handleApply = () => {
  emit('apply', props.job.id);
};

watch(() => props.job, () => {
  selectedImageIndex.value = 0;
});
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
  to { transform: translateX(0); }
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

.side-company-type {
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
}

.side-meta-row + .side-meta-row {
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

.side-action-buttons {
  display: flex;
  gap: 12px;
  margin: 20px 0 10px 0;
  flex-wrap: wrap;
}

.btn-apply,
.btn-save {
  flex: 1;
  min-width: 160px;
  padding: 13px 20px;
  border: none;
  border-radius: 12px;
  font-size: 15px;
  font-weight: 600;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 8px;
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.08);
  letter-spacing: 0.02em;
}

.btn-apply {
  background: linear-gradient(135deg, #037266 0%, #026b5b 100%);
  color: #ffffff;
  border: 2px solid transparent;
}

.btn-apply:hover:not(:disabled) {
  background: linear-gradient(135deg, #026b5b 0%, #025a4d 100%);
  transform: translateY(-2px);
  box-shadow: 0 6px 20px rgba(3, 114, 102, 0.25);
}

.btn-apply.applied {
  background: linear-gradient(135deg, #10b981 0%, #059669 100%);
  cursor: default;
  animation: successPulse 0.6s ease;
}

.btn-apply:disabled {
  cursor: not-allowed;
  opacity: 0.9;
}

@keyframes successPulse {
  0%, 100% { transform: scale(1); }
  50% { transform: scale(1.05); }
}

.btn-save {
  background: #ffffff;
  color: #037266;
  border: 2px solid #cee5e2;
}

.btn-save:hover {
  background: #f0f8f7;
  border-color: #037266;
  transform: translateY(-2px);
  box-shadow: 0 6px 20px rgba(3, 114, 102, 0.15);
}

.btn-save.saved {
  background: linear-gradient(135deg, #ffd43b 0%, #ffc800 100%);
  color: #ffffff;
  border-color: #ffd43b;
  animation: savePulse 0.5s ease;
}

.btn-save.saved:hover {
  background: linear-gradient(135deg, #ffc800 0%, #ffb700 100%);
  box-shadow: 0 6px 20px rgba(255, 212, 59, 0.3);
}

@keyframes savePulse {
  0%, 100% { transform: scale(1); }
  50% { transform: scale(1.05); }
}

.btn-apply svg,
.btn-save svg {
  transition: transform 0.3s ease;
}

.btn-apply:hover:not(:disabled) svg,
.btn-save:hover svg {
  transform: scale(1.1);
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
  white-space: pre-line;
}

@media (max-width: 900px) {
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

  .side-action-buttons {
    flex-direction: column;
  }

  .btn-apply,
  .btn-save {
    min-width: 100%;
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