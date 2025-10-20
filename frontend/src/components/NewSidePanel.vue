<!-- frontend/src/components/NewsSidePanel.vue -->
<template>
  <div v-if="news" class="side-panel-overlay" @click.self="$emit('close')">
    <div class="side-panel">
      <button class="side-close" @click="$emit('close')">&times;</button>
      <div class="side-panel-inner">
        <div class="side-header">
          <h2>{{ news.title }}</h2>

          <!-- แสดง category และ status (ถ้ามี) -->
          <div v-if="news.category || (showStatus && news.status)" class="news-meta-inline">
            <span v-if="news.category" class="category-text-meta">{{ news.category }}</span>
            <span v-if="showStatus && news.status" class="status-badge" :class="news.status">
              <span class="status-dot"></span>
              {{ getStatusText(news.status) }}
            </span>
          </div>

          <p class="news-date">
            <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" width="14" height="14" fill="none"
              stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
              <rect x="3" y="4" width="18" height="18" rx="2" ry="2" />
              <line x1="16" y1="2" x2="16" y2="6" />
              <line x1="8" y1="2" x2="8" y2="6" />
              <line x1="3" y1="10" x2="21" y2="10" />
            </svg>
            {{ formatDate(news.publishDate || news.date) }}
          </p>
        </div>

        <!-- Slot สำหรับปุ่มกำหนดเอง (เช่น ปุ่มแก้ไข/ลบสำหรับอาจารย์) -->
        <slot name="actions">
          <!-- Default action buttons สำหรับนักศึกษา -->
          <div v-if="showActions" class="side-action-buttons">
            <button class="btn-save" :class="{ saved: isSaved }" @click="$emit('toggle-save', news.id)">
              <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24"
                :fill="isSaved ? 'currentColor' : 'none'" stroke="currentColor" stroke-width="2" stroke-linecap="round"
                stroke-linejoin="round">
                <path d="M19 21l-7-5-7 5V5a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2z"></path>
              </svg>
              <span>{{ isSaved ? 'บันทึกแล้ว' : 'บันทึก' }}</span>
            </button>
          </div>
        </slot>

        <p class="side-description" v-html="news.description"></p>

        <div v-if="news.images?.length" class="side-images">
          <img v-for="(img, i) in news.images" :key="i" :src="img" :alt="`Image ${i + 1}`" />
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
const props = defineProps({
  news: {
    type: Object,
    default: null
  },
  showActions: {
    type: Boolean,
    default: true
  },
  isSaved: {
    type: Boolean,
    default: false
  },
  showStatus: {
    type: Boolean,
    default: false
  }
});

defineEmits(['close', 'toggle-save']);

const formatDate = (dateStr) => {
  const d = new Date(dateStr);
  const monthNames = [
    "มกราคม", "กุมภาพันธ์", "มีนาคม", "เมษายน", "พฤษภาคม", "มิถุนายน",
    "กรกฎาคม", "สิงหาคม", "กันยายน", "ตุลาคม", "พฤศจิกายน", "ธันวาคม",
  ];
  return `${d.getDate()} ${monthNames[d.getMonth()]} ${d.getFullYear() + 543}`;
};

const getStatusText = (status) => {
  const statusMap = {
    published: 'เผยแพร่แล้ว',
    scheduled: 'รอเผยแพร่',
  };
  return statusMap[status] || status;
};
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
  padding: 46px 34px 38px 34px;
  display: flex;
  flex-direction: column;
  gap: 18px;
}

.side-header h2 {
  font-size: 24px;
  font-weight: 700;
  margin-bottom: 12px;
  color: #222;
  padding-right: 40px;
  line-height: 1.3;
}

.news-meta-inline {
  display: flex;
  gap: 12px;
  margin-bottom: 12px;
  align-items: center;
  flex-wrap: wrap;
}

.category-text-meta {
  font-size: 13px;
  color: #666;
  font-weight: 500;
}

.status-badge {
  display: inline-flex;
  align-items: center;
  gap: 6px;
  padding: 6px 12px;
  border-radius: 20px;
  font-size: 13px;
  font-weight: 600;
  white-space: nowrap;
}

.status-badge.published {
  background: #d1fae5;
  color: #065f46;
}

.status-badge.scheduled {
  background: #fef3c7;
  color: #92400e;
}

.status-dot {
  width: 8px;
  height: 8px;
  border-radius: 50%;
  background: currentColor;
}

.news-date {
  display: inline-flex;
  align-items: center;
  gap: 6px;
  font-size: 13px;
  color: #555;
  margin: 0;
}

.news-date svg {
  flex-shrink: 0;
}

.side-action-buttons {
  display: flex;
  gap: 12px;
  margin: 20px 0 10px 0;
  flex-wrap: wrap;
}

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

.btn-save:active {
  transform: translateY(0);
  box-shadow: 0 2px 8px rgba(3, 114, 102, 0.1);
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

  0%,
  100% {
    transform: scale(1);
  }

  50% {
    transform: scale(1.05);
  }
}

.btn-save svg {
  transition: transform 0.3s ease;
}

.btn-save:hover svg {
  transform: scale(1.1);
}

.side-description {
  margin-top: 15px;
  font-size: 15px;
  line-height: 1.7;
  color: #444;
  white-space: pre-line;
}

.side-images {
  display: flex;
  flex-wrap: wrap;
  gap: 12px;
  margin-top: 20px;
}

.side-images img {
  width: calc(50% - 6px);
  height: 150px;
  object-fit: cover;
  border-radius: 10px;
  transition: transform 0.25s ease;
}

.side-images img:hover {
  transform: scale(1.05);
}

@media (max-width: 900px) {
  .side-panel-inner {
    padding: 28px 20px 21px 20px;
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

  .side-panel-inner {
    padding: 14px 12px 10px 12px;
  }

  .side-action-buttons {
    flex-direction: column;
  }

  .btn-save {
    min-width: 100%;
  }

  .side-images img {
    width: 100%;
  }
}
</style>