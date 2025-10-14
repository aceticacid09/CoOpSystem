<!-- frontend/src/components/FavoriteSections.vue -->
<template>
  <section class="favorites-section container">
    <!-- Tabs -->
    <div class="tabs">
      <button v-for="tab in tabs" :key="tab" class="tab-btn" :class="{ active: activeTab === tab }"
        @click="selectTab(tab)">
        {{ tab }}
      </button>
    </div>

    <!-- Content for ข่าวสารที่บันทึก -->
    <div v-if="activeTab === 'ข่าวสารที่บันทึก'" class="favorites-content">
      <div v-if="savedNews.length === 0" class="empty-state">
        <svg xmlns="http://www.w3.org/2000/svg" width="64" height="64" viewBox="0 0 24 24" fill="none"
          stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
          <path d="M19 21l-7-5-7 5V5a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2z"></path>
        </svg>
        <p>ยังไม่มีข่าวสารที่บันทึก</p>
      </div>

      <div v-else class="news-list">
        <div v-for="item in savedNews" :key="item.id" class="news-card" @click="openNewsModal(item)">
          <!-- Bookmark Button -->
          <button class="bookmark-btn bookmarked" @click="removeNewsBookmark(item.id, $event)"
            aria-label="ยกเลิกบันทึก">
            <svg xmlns="http://www.w3.org/2000/svg" width="22" height="22" viewBox="0 0 24 24" fill="#ffd43b"
              stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
              <path d="M19 21l-7-5-7 5V5a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2z"></path>
            </svg>
          </button>

          <!-- IMAGE -->
          <div class="news-img" v-if="item.images?.length">
            <img v-if="item.images.length === 1" :src="item.images[0]" alt="news" />
            <div v-else class="image-grid" :class="`images-${Math.min(item.images.length, 4)}`">
              <div v-for="(img, i) in item.images.slice(0, 4)" :key="i" class="image-cell">
                <img :src="img" alt="news" />
                <div v-if="i === 3 && item.images.length > 4" class="image-overlay">
                  +{{ item.images.length - 4 }}
                </div>
              </div>
            </div>
          </div>

          <!-- CONTENT -->
          <div class="news-content">
            <h3 class="news-title">{{ item.title }}</h3>
            <p class="news-date">
              <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" width="14" height="14" fill="none"
                stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                <rect x="3" y="4" width="18" height="18" rx="2" ry="2" />
                <line x1="16" y1="2" x2="16" y2="6" />
                <line x1="8" y1="2" x2="8" y2="6" />
                <line x1="3" y1="10" x2="21" y2="10" />
              </svg>
              ประกาศเมื่อ {{ formatDate(item.date) }}
            </p>
            <p class="news-desc">{{ item.description }}</p>
          </div>
        </div>
      </div>
    </div>

    <!-- Content for งานที่บันทึก -->
    <div v-else-if="activeTab === 'งานที่บันทึก'" class="favorites-content">
      <div v-if="savedJobs.length === 0" class="empty-state">
        <svg xmlns="http://www.w3.org/2000/svg" width="64" height="64" viewBox="0 0 24 24" fill="none"
          stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
          <path d="M19 21l-7-5-7 5V5a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2z"></path>
        </svg>
        <p>ยังไม่มีงานที่บันทึก</p>
      </div>

      <div v-else class="jobs-list">
        <div v-for="job in savedJobs" :key="job.id" class="job-card" @click="openJobPanel(job)">
          <!-- Bookmark Button -->
          <button class="bookmark-btn bookmarked" @click="removeJobBookmark(job.id, $event)" aria-label="ยกเลิกบันทึก">
            <svg xmlns="http://www.w3.org/2000/svg" width="22" height="22" viewBox="0 0 24 24" fill="#ffd43b"
              stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
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
                  <span>ปิดรับสมัคร {{ job.closeDate }}</span>
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
      </div>
    </div>

    <!-- ========== NEWS DETAIL SIDE PANEL ========== -->
    <div v-if="selectedNews" class="side-panel-overlay" @click.self="closeNewsModal">
      <div class="side-panel">
        <button class="side-close" @click="closeNewsModal">&times;</button>
        <div class="side-panel-inner">
          <div class="side-header">
            <h2>{{ selectedNews.title }}</h2>
            <p class="news-date">
              <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" width="14" height="14" fill="none"
                stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                <rect x="3" y="4" width="18" height="18" rx="2" ry="2" />
                <line x1="16" y1="2" x2="16" y2="6" />
                <line x1="8" y1="2" x2="8" y2="6" />
                <line x1="3" y1="10" x2="21" y2="10" />
              </svg>
              {{ formatDate(selectedNews.date) }}
            </p>
          </div>
          <!-- Action Button - แสดงเฉพาะเมื่อ showActions เป็น true -->
            <div v-if="showActions" class="side-action-buttons">
              <button class="btn-save" :class="{ saved: isSavedNews(selectedNews.id) }"
                @click="toggleSaveNews(selectedNews.id)">
                <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24"
                  :fill="isSavedNews(selectedNews.id) ? 'currentColor' : 'none'" stroke="currentColor" stroke-width="2"
                  stroke-linecap="round" stroke-linejoin="round">
                  <path d="M19 21l-7-5-7 5V5a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2z"></path>
                </svg>
                <span>{{ isSavedNews(selectedNews.id) ? 'บันทึกแล้ว' : 'บันทึก' }}</span>
              </button>
            </div>
          <p class="side-description">{{ selectedNews.description }}</p>
          <div v-if="selectedNews.images?.length" class="side-images">
            <img v-for="(img, i) in selectedNews.images" :key="i" :src="img" alt="news" />
          </div>
        </div>
      </div>
    </div>

    <!-- ========== JOB DETAIL SIDE PANEL ========== -->
    <div v-if="selectedJob" class="side-panel-overlay" @click.self="closeJobPanel">
      <div class="side-panel">
        <button class="side-close" @click="closeJobPanel">&times;</button>
        <div class="side-panel-inner">
          <div class="side-header-top">
            <img class="side-logo" :src="selectedJob.logo" alt="logo" />
            <div class="side-title-block">
              <div class="side-title">{{ selectedJob.title }}</div>
              <div class="side-company-row">
                <span class="side-company-name">{{ selectedJob.company }}</span>
                <span class="side-company-sep"></span>
                <span class="side-company-type">ดูรายละเอียดงาน</span>
              </div>
              <span class="side-department-label">{{ selectedJob.department }}</span>
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
              ปิดรับสมัคร {{ selectedJob.closeDate }}
            </span>
          </div>
          <!-- Action Buttons - แสดงเฉพาะเมื่อ showActions เป็น true -->
<div v-if="showActions" class="side-action-buttons">
  <button class="btn-apply" :class="{ applied: isApplied(selectedJob.id) }" @click="applyJob(selectedJob.id)"
    :disabled="isApplied(selectedJob.id)">
    <svg v-if="!isApplied(selectedJob.id)" xmlns="http://www.w3.org/2000/svg" width="20" height="20"
      viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round"
      stroke-linejoin="round">
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
    <span>{{ isApplied(selectedJob.id) ? 'สมัครงานแล้ว' : 'สมัครงาน' }}</span>
  </button>

  <button class="btn-save" :class="{ saved: isSavedJob(selectedJob.id) }" @click="toggleSaveJob(selectedJob.id)">
    <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24"
      :fill="isSavedJob(selectedJob.id) ? 'currentColor' : 'none'" stroke="currentColor" stroke-width="2"
      stroke-linecap="round" stroke-linejoin="round">
      <path d="M19 21l-7-5-7 5V5a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2z"></path>
    </svg>
    <span>{{ isSavedJob(selectedJob.id) ? 'บันทึกแล้ว' : 'บันทึก' }}</span>
  </button>
</div>
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
  </section>
</template>

<script setup>
import { ref, computed, watch } from "vue";
import { newsList } from "../data/newsData.js";
import jobsList from "../data/jobsData.js";

// Props
const props = defineProps({
  savedNewsIds: {
    type: Array,
    default: () => []
  },
  savedJobsIds: {
    type: Array,
    default: () => []
  },
  showBookmark: {
    type: Boolean,
    default: true  // แสดงปุ่ม bookmark ใน Favorites
  },
  showActions: {
    type: Boolean,
    default: true  // แสดงปุ่ม Action ใน Side Panel
  }
});

// Applied and Saved state
const appliedJobs = ref(new Set());
const savedNewsInPanel = ref(new Set(props.savedNewsIds));
const savedJobsInPanel = ref(new Set(props.savedJobsIds));

// Methods for actions
const applyJob = (jobId) => {
  appliedJobs.value.add(jobId);
  appliedJobs.value = new Set(appliedJobs.value);
};

const toggleSaveJob = (jobId) => {
  if (savedJobsInPanel.value.has(jobId)) {
    savedJobsInPanel.value.delete(jobId);
  } else {
    savedJobsInPanel.value.add(jobId);
  }
  savedJobsInPanel.value = new Set(savedJobsInPanel.value);
};

const toggleSaveNews = (newsId) => {
  if (savedNewsInPanel.value.has(newsId)) {
    savedNewsInPanel.value.delete(newsId);
  } else {
    savedNewsInPanel.value.add(newsId);
  }
  savedNewsInPanel.value = new Set(savedNewsInPanel.value);
};

const isApplied = (jobId) => appliedJobs.value.has(jobId);
const isSavedJob = (jobId) => savedJobsInPanel.value.has(jobId);
const isSavedNews = (newsId) => savedNewsInPanel.value.has(newsId);

// Emit events
const emit = defineEmits(['remove-news-bookmark', 'remove-job-bookmark']);

// Tabs
const tabs = ["ข่าวสารที่บันทึก", "งานที่บันทึก"];
const activeTab = ref(tabs[0]);

// Selected items for modals
const selectedNews = ref(null);
const selectedJob = ref(null);
const selectedImageIndex = ref(0);

// Computed: filter saved items
const savedNews = computed(() => {
  return newsList.filter(news => props.savedNewsIds.includes(news.id));
});

const savedJobs = computed(() => {
  return jobsList.filter(job => props.savedJobsIds.includes(job.id));
});

// Methods
const selectTab = (tab) => {
  activeTab.value = tab;
};

const openNewsModal = (item) => {
  selectedNews.value = { ...item };
};

const closeNewsModal = () => {
  selectedNews.value = null;
};

const openJobPanel = (job) => {
  selectedJob.value = job;
  selectedImageIndex.value = 0;
};

const closeJobPanel = () => {
  selectedJob.value = null;
};

const removeNewsBookmark = (newsId, event) => {
  event.stopPropagation();
  emit('remove-news-bookmark', newsId);
};

const removeJobBookmark = (jobId, event) => {
  event.stopPropagation();
  emit('remove-job-bookmark', jobId);
};

const truncateText = (text, len) => {
  if (!text) return "";
  return text.length <= len ? text : text.slice(0, len);
};

const formatDate = (dateStr) => {
  const d = new Date(dateStr);
  const monthNames = [
    "มกราคม", "กุมภาพันธ์", "มีนาคม", "เมษายน", "พฤษภาคม", "มิถุนายน",
    "กรกฎาคม", "สิงหาคม", "กันยายน", "ตุลาคม", "พฤศจิกายน", "ธันวาคม",
  ];
  return `${d.getDate()} ${monthNames[d.getMonth()]} ${d.getFullYear() + 543}`;
};

// Watch for job selection to reset image index
watch(selectedJob, () => {
  selectedImageIndex.value = 0;
});
</script>

<style scoped>
@import url("https://fonts.googleapis.com/css2?family=Kanit:wght@400;600;700&display=swap");

* {
  font-family: "Kanit", "Prompt", "Sarabun", "Noto Sans Thai", sans-serif;
  box-sizing: border-box;
}

.favorites-section {
  padding: 0;
}

/* ========== TABS ========== */
.tabs {
  display: flex;
  gap: 25px;
  border-bottom: 1px solid #eee;
  margin-bottom: 20px;
  flex-wrap: wrap;
}

.tab-btn {
  background: none;
  border: none;
  font-size: 16px;
  color: #555;
  padding-bottom: 8px;
  cursor: pointer;
  position: relative;
  transition: 0.2s;
}

.tab-btn.active {
  color: #037266;
  font-weight: 600;
}

.tab-btn.active::after {
  content: "";
  position: absolute;
  left: 0;
  bottom: -1px;
  width: 100%;
  height: 2px;
  background: #037266;
}

/* ========== CONTENT ========== */
.favorites-content {
  padding: 20px 0;
}

/* ========== EMPTY STATE ========== */
.empty-state {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  padding: 60px 20px;
  color: #999;
}

.empty-state svg {
  margin-bottom: 20px;
  opacity: 0.3;
}

.empty-state p {
  font-size: 16px;
  font-weight: 500;
}

/* ========== NEWS LIST ========== */
.news-list {
  display: flex;
  flex-direction: column;
  gap: 35px;
}

.news-card {
  position: relative;
  display: flex;
  gap: 20px;
  padding: 20px;
  flex-wrap: wrap;
  cursor: pointer;
  transition: transform 0.2s ease, box-shadow 0.2s ease;
  background: #fff;
  border-radius: 16px;
  box-shadow: 0 2px 12px rgba(3, 114, 102, 0.07);
  border: 1px solid #eef0ef;
}

.news-card:hover {
  transform: translateY(-3px);
  box-shadow: 0 6px 24px rgba(3, 114, 102, 0.14);
}

.news-img img {
  width: 220px;
  height: 160px;
  object-fit: cover;
  border-radius: 8px;
}

.news-content {
  flex: 1;
}

.news-title {
  font-size: 18px;
  color: #000;
  font-weight: 600;
  margin: 0;
}

.news-date {
  display: inline-flex;
  align-items: center;
  gap: 6px;
  font-size: 13px;
  color: #555;
  margin-top: 6px;
  margin-bottom: 8px;
}

.news-date svg {
  flex-shrink: 0;
}

.news-desc {
  font-size: 14px;
  color: #444;
  line-height: 1.6;
  margin: 0;
}

/* ========== JOBS LIST ========== */
.jobs-list {
  display: flex;
  flex-direction: column;
  gap: 32px;
}

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

/* ========== IMAGE GRID ========== */
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

.image-grid.images-3 .image-cell:nth-child(1) {
  grid-area: a;
}

.image-grid.images-3 .image-cell:nth-child(2) {
  grid-area: b;
}

.image-grid.images-3 .image-cell:nth-child(3) {
  grid-area: c;
}

.image-grid.images-4 {
  grid-template-columns: repeat(2, 1fr);
  grid-template-rows: repeat(2, 80px);
  width: 220px;
}

/* ========== BOOKMARK BUTTON ========== */
.bookmark-btn {
  position: absolute;
  top: 18px;
  right: 18px;
  width: 42px;
  height: 42px;
  border: none;
  border-radius: 50%;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  box-shadow: 0 2px 8px rgba(3, 114, 102, 0.08);
  transition: all 0.25s cubic-bezier(0.4, 0, 0.2, 1);
  z-index: 10;
  border: 1.5px solid #ffd43b;
  background: linear-gradient(135deg, #ffd43b 0%, #ffc800 100%);
  color: #ffffff;
}

.bookmark-btn:hover {
  background: linear-gradient(135deg, #ffc800 0%, #ffb700 100%);
  box-shadow: 0 4px 20px rgba(255, 212, 59, 0.4);
  transform: translateY(-2px) scale(1.05);
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
}

.side-header h2 {
  font-size: 24px;
  font-weight: 700;
  margin-bottom: 8px;
  color: #222;
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

/* ========== RESPONSIVE ========== */
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
  .bookmark-btn {
    width: 38px;
    height: 38px;
    top: 14px;
    right: 14px;
  }

  .bookmark-btn svg {
    width: 19px;
    height: 19px;
  }

  .news-card,
  .job-card {
    padding: 14px;
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

  .side-images img {
    width: 100%;
  }
}

/* ========== ACTION BUTTONS (in side panel) ========== */
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

@media (max-width: 650px) {
  .side-action-buttons {
    flex-direction: column;
  }

  .btn-apply,
  .btn-save {
    min-width: 100%;
  }
}
</style>