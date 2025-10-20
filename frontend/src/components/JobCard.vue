<!-- frontend/src/components/JobCard.vue -->
<template>
  <div class="job-card" @click="$emit('click', job)">
    <!-- Bookmark Button -->
    <button v-if="showBookmark" class="bookmark-btn" :class="{ bookmarked: isBookmarked }"
      @click="handleBookmarkClick" :aria-label="isBookmarked ? 'ยกเลิกบันทึก' : 'บันทึกงาน'">
      <svg xmlns="http://www.w3.org/2000/svg" width="22" height="22" viewBox="0 0 24 24"
        :fill="isBookmarked ? '#ffd43b' : 'none'" stroke="currentColor" stroke-width="2" stroke-linecap="round"
        stroke-linejoin="round">
        <path d="M19 21l-7-5-7 5V5a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2z"></path>
      </svg>
    </button>

    <div class="job-header">
      <img class="company-logo" :src="job.logo" alt="logo" />
      <div class="job-header-info">
        <div class="company-name">{{ job.company }}</div>
        <div class="department-label">{{ job.department }}</div>
        <div class="job-meta">
          <div class="job-meta-row">
            <span class="meta-icon">
              <svg width="16" height="16" fill="none" stroke="#9CA3AF" stroke-width="2" stroke-linecap="round"
                stroke-linejoin="round" viewBox="0 0 24 24">
                <rect x="3" y="4" width="18" height="18" rx="2" ry="2" />
                <line x1="16" y1="2" x2="16" y2="6" />
                <line x1="8" y1="2" x2="8" y2="6" />
                <line x1="3" y1="10" x2="21" y2="10" />
              </svg>
            </span>
            <span>ประกาศเมื่อ {{ job.postDate }}</span>
          </div>
          <div class="job-meta-row">
            <span class="meta-icon">
              <svg width="16" height="16" fill="none" stroke="#9CA3AF" stroke-width="2" stroke-linecap="round"
                stroke-linejoin="round" viewBox="0 0 24 24">
                <circle cx="12" cy="12" r="10"></circle>
                <polyline points="12 6 12 12 16 14"></polyline>
              </svg>
            </span>
            <span>ปิดรับสมัคร {{ job.closeDate }} <span v-if="job.applyDuration">(ระยะเวลา {{ job.applyDuration
                }})</span></span>
          </div>
        </div>
      </div>
    </div>

    <div class="job-content">
      <div class="job-images" v-if="job.images && job.images.length">
        <img v-if="job.images.length === 1" :src="job.images[0]" class="job-img" alt="job" />
        <div v-else class="image-grid" :class="`images-${Math.min(job.images.length, 4)}`">
          <div v-for="(img, i) in job.images.slice(0, 4)" :key="i" class="image-cell">
            <img :src="img" alt="job" />
            <div v-if="i === 3 && job.images.length > 4" class="image-overlay">
              +{{ job.images.length - 4 }}
            </div>
          </div>
        </div>
      </div>
      <div class="job-details">
        <div class="job-title">{{ job.title }}</div>
        <div class="job-details-title">รายละเอียดงาน</div>
        <div class="job-description">
          {{ truncateText(job.description, 160) }}
          <span v-if="job.description.length > 160" class="show-more">... ดูเพิ่มเติม</span>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
const props = defineProps({
  job: {
    type: Object,
    required: true
  },
  showBookmark: {
    type: Boolean,
    default: false
  },
  isBookmarked: {
    type: Boolean,
    default: false
  }
});

const emit = defineEmits(['click', 'toggle-bookmark']);

const handleBookmarkClick = (event) => {
  event.stopPropagation();
  emit('toggle-bookmark', props.job.id);
};

const truncateText = (text, len) => {
  if (!text) return "";
  return text.length <= len ? text : text.slice(0, len);
};
</script>

<style scoped>
.job-card {
  position: relative;
  background: #fff;
  border-radius: 16px;
  box-shadow: 0 2px 12px rgba(3, 114, 102, 0.07);
  border: 1px solid #eef0ef;
  padding: 25px 30px 22px 30px;
  display: flex;
  flex-direction: column;
  gap: 16px;
  transition: box-shadow 0.2s, transform 0.2s;
  cursor: pointer;
}

.job-card:hover {
  box-shadow: 0 6px 24px rgba(3, 114, 102, 0.14);
  transform: translateY(-2px);
}

.bookmark-btn {
  position: absolute;
  top: 22px;
  right: 26px;
  width: 42px;
  height: 42px;
  border: none;
  border-radius: 50%;
  background: #ffffff;
  color: #9CA3AF;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  box-shadow: 0 2px 8px rgba(3, 114, 102, 0.08);
  transition: all 0.25s cubic-bezier(0.4, 0, 0.2, 1);
  z-index: 10;
  border: 1.5px solid #f0f0f0;
}

.bookmark-btn:hover {
  background: #f8f9fa;
  transform: translateY(-2px) scale(1.05);
  box-shadow: 0 4px 16px rgba(3, 114, 102, 0.15);
  border-color: #e0e0e0;
}

.bookmark-btn.bookmarked {
  background: linear-gradient(135deg, #ffd43b 0%, #ffc800 100%);
  color: #ffffff;
  border-color: #ffd43b;
  animation: bookmarkPulse 0.5s ease;
}

.bookmark-btn.bookmarked:hover {
  background: linear-gradient(135deg, #ffc800 0%, #ffb700 100%);
  box-shadow: 0 4px 20px rgba(255, 212, 59, 0.4);
}

@keyframes bookmarkPulse {
  0%, 100% { transform: scale(1); }
  50% { transform: scale(1.15); }
}

.job-header {
  display: flex;
  align-items: flex-start;
  gap: 20px;
}

.company-logo {
  width: 58px;
  height: 58px;
  border-radius: 50%;
  object-fit: cover;
  background: #f0f0f0;
  border: 1.5px solid #e1e1e1;
  flex-shrink: 0;
}

.job-header-info {
  display: flex;
  flex-direction: column;
  gap: 6px;
  flex: 1;
}

.company-name {
  font-size: 17px;
  font-weight: 700;
  color: #333;
  text-transform: uppercase;
  letter-spacing: 0.03em;
}

.department-label {
  display: inline-block;
  background-color: #cee5e2;
  color: #037266;
  border-radius: 7px;
  font-size: 13px;
  font-weight: 500;
  padding: 3px 14px;
  letter-spacing: 0.04em;
  align-self: flex-start;
}

.job-meta {
  font-size: 13px;
  color: #6b6b6b;
  font-weight: 500;
  display: flex;
  flex-direction: column;
  gap: 3px;
  margin-top: 2px;
}

.job-meta-row {
  display: flex;
  align-items: center;
  gap: 5px;
}

.meta-icon {
  display: flex;
  align-items: center;
  justify-content: center;
  margin-right: 3px;
}

.job-content {
  display: flex;
  gap: 28px;
  align-items: flex-start;
  margin-left: 75px;
}

.job-images {
  min-width: 220px;
  max-width: 220px;
  flex-shrink: 0;
}

.job-img {
  width: 220px;
  height: 160px;
  object-fit: cover;
  border-radius: 8px;
}

.image-grid {
  display: grid;
  gap: 6px;
}

.image-cell {
  position: relative;
}

.image-cell img {
  width: 100%;
  height: 100%;
  border-radius: 8px;
  object-fit: cover;
}

.image-overlay {
  position: absolute;
  inset: 0;
  background: rgba(0, 0, 0, 0.6);
  color: #fff;
  font-size: 18px;
  font-weight: bold;
  display: flex;
  align-items: center;
  justify-content: center;
  border-radius: 8px;
}

.image-grid.images-2 {
  grid-template-columns: repeat(2, 1fr);
  grid-auto-rows: 160px;
  width: 220px;
}

.image-grid.images-3 {
  grid-template-areas: "a b" "c c";
  grid-template-columns: repeat(2, 1fr);
  grid-template-rows: 80px 80px;
  width: 220px;
}

.image-grid.images-3 .image-cell:nth-child(1) { grid-area: a; }
.image-grid.images-3 .image-cell:nth-child(2) { grid-area: b; }
.image-grid.images-3 .image-cell:nth-child(3) { grid-area: c; }

.image-grid.images-4 {
  grid-template-columns: repeat(2, 1fr);
  grid-template-rows: repeat(2, 80px);
  width: 220px;
}

.job-details {
  flex: 1;
  display: flex;
  flex-direction: column;
  gap: 8px;
  min-width: 0;
}

.job-title {
  font-size: 22px;
  font-weight: 700;
  color: #222;
  margin: 0;
}

.job-details-title {
  font-size: 15px;
  font-weight: 600;
  color: #777;
  margin: 0;
}

.job-description {
  font-size: 14px;
  color: #444;
  line-height: 1.6;
  font-weight: 400;
  white-space: pre-line;
  margin: 0;
}

.show-more {
  color: #009688;
  font-weight: 600;
  font-size: 14px;
  cursor: pointer;
  margin-left: 4px;
}

@media (max-width: 900px) {
  .job-content {
    flex-direction: column;
    gap: 20px;
    margin-left: 0;
  }

  .job-images {
    max-width: 100%;
    min-width: 100%;
  }

  .job-img {
    width: 100%;
  }

  .image-grid.images-2,
  .image-grid.images-3,
  .image-grid.images-4 {
    width: 100%;
  }
}

@media (max-width: 650px) {
  .bookmark-btn {
    width: 38px;
    height: 38px;
    top: 16px;
    right: 16px;
  }

  .bookmark-btn svg {
    width: 19px;
    height: 19px;
  }

  .job-card {
    padding: 18px 16px;
  }

  .company-logo {
    width: 46px;
    height: 46px;
  }

  .job-header {
    gap: 14px;
  }

  .job-title {
    font-size: 18px;
  }
}
</style>