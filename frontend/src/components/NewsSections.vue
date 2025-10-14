<!-- frontend/src/components/NewsSections.vue -->
<template>
  <!-- ========== NEWS & EVENTS SECTION ========== -->
  <section class="news-events container" :style="{ '--news-max-height': maxHeight || 'none' }">

    <!-- ========== TAB NAVIGATION ========== -->
    <div class="tabs">
      <button v-for="tab in tabs" :key="tab" class="tab-btn" :class="{ active: activeTab === tab }"
        @click="activeTab = tab">
        {{ tab }}
      </button>
    </div>

    <!-- ========== SEARCH & SORT ========== -->
    <div class="search-sort">
      <!-- SEARCH -->
      <div class="search-box">
        <input type="text" v-model="searchText" placeholder="ค้นหาหัวข้อข่าว" />
      </div>
      <button class="btn-search">ค้นหา</button>

      <!-- SORT -->
      <button @click="toggleSort" class="btn-sort" :title="isAscending ? 'เรียงจากเก่าสุด' : 'เรียงจากใหม่สุด'">
        <svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" fill="none" stroke="#9CA3AF" stroke-width="2"
          stroke-linecap="round" stroke-linejoin="round">
          <path d="M3 6h18M3 12h12M3 18h6" />
        </svg>
      </button>
    </div>

    <!-- ========== NEWS LIST ========== -->
    <div class="news-list">
      <div v-for="(item, index) in filteredNews" :key="index" class="news-card" @click="openModal(item, index)">
        
        <!-- Bookmark Button - แสดงเฉพาะเมื่อ showBookmark เป็น true -->
        <button v-if="showBookmark" class="bookmark-btn" :class="{ bookmarked: isBookmarked(item.id || index) }"
          @click="toggleBookmark(item.id || index, $event)" :aria-label="isBookmarked(item.id || index) ? 'ยกเลิกบันทึก' : 'บันทึกข่าว'">
          <svg xmlns="http://www.w3.org/2000/svg" width="22" height="22" viewBox="0 0 24 24"
            :fill="isBookmarked(item.id) ? '#ffd43b' : 'none'" stroke="currentColor" stroke-width="2"
            stroke-linecap="round" stroke-linejoin="round">
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

    <!-- ========== SIDE PANEL (NEWS DETAIL) ========== -->
    <div v-if="selectedNews" class="side-panel-overlay" @click.self="closeModal">
      <div class="side-panel">
        <!-- CLOSE BUTTON -->
        <button class="side-close" @click="closeModal">&times;</button>

        <div class="side-panel-inner">
          <!-- HEADER -->
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

          <!-- Action Button - แสดงเฉพาะเมื่อ showBookmark เป็น true -->
          <div v-if="showBookmark" class="side-action-buttons">
            <button class="btn-save" :class="{ saved: isSaved(selectedNews.uniqueId) }" @click="toggleSaveNews(selectedNews.uniqueId)">
              <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24"
                :fill="isSaved(selectedNews.uniqueId) ? 'currentColor' : 'none'" stroke="currentColor" stroke-width="2"
                stroke-linecap="round" stroke-linejoin="round">
                <path d="M19 21l-7-5-7 5V5a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2z"></path>
              </svg>
              <span>{{ isSaved(selectedNews.uniqueId) ? 'บันทึกแล้ว' : 'บันทึก' }}</span>
            </button>
          </div>

          <!-- DESCRIPTION -->
          <p class="side-description">{{ selectedNews.description }}</p>

          <!-- IMAGES -->
          <div v-if="selectedNews.images?.length" class="side-images">
            <img v-for="(img, i) in selectedNews.images" :key="i" :src="img" alt="news" />
          </div>
        </div>
      </div>
    </div>
  </section>
</template>

<script setup>
/* ========== IMPORTS ========== */
import { ref, computed } from "vue";
import { newsList } from "../data/newsData.js";

/* ========== PROPS ========== */
const props = defineProps({
  items: { type: Array, default: () => newsList },
  maxHeight: { type: String, default: null },
  showBookmark: {
    type: Boolean,
    default: false
  },
  initialBookmarked: {
    type: Array,
    default: () => [] // เช่น [1, 3, 5] = ข่าว id 1,3,5 ถูกบันทึกไว้แล้ว
  },
  initialSaved: {
    type: Array,
    default: () => [] // เช่น [2, 4] = ข่าว id 2,4 ถูกบันทึกไว้แล้ว
  }
});

/* ========== STATE ========== */
const tabs = ["ข่าวทั้งหมด", "ประชาสัมพันธ์", "ข่าวด่วน", "กิจกรรม", "ประกาศผลการคัดเลือก"];
const activeTab = ref(tabs[0]);
const searchText = ref("");
const isAscending = ref(true);
const selectedNews = ref(null);

// Bookmark state - เริ่มต้นจาก props
const bookmarkedNews = ref(new Set(props.initialBookmarked));
const savedNews = ref(new Set(props.initialSaved));

/* ========== COMPUTED (FILTER & SORT) ========== */
const filteredNews = computed(() => {
  let filtered = props.items.filter(
    n =>
      (activeTab.value === "ข่าวทั้งหมด" || n.category === activeTab.value) &&
      (n.title.includes(searchText.value) || n.description.includes(searchText.value))
  );
  return isAscending.value ? filtered : filtered.slice().reverse();
});

/* ========== METHODS ========== */
const toggleSort = () => (isAscending.value = !isAscending.value);
const openModal = (item, index) => {
  selectedNews.value = { ...item, uniqueId: item.id || index };
};
const closeModal = () => (selectedNews.value = null);

// Bookmark functions
const toggleBookmark = (newsId, event) => {
  event.stopPropagation();
  if (bookmarkedNews.value.has(newsId)) {
    bookmarkedNews.value.delete(newsId);
  } else {
    bookmarkedNews.value.add(newsId);
  }
  bookmarkedNews.value = new Set(bookmarkedNews.value);
};

const isBookmarked = (newsId) => {
  return bookmarkedNews.value.has(newsId);
};

const toggleSaveNews = (newsId) => {
  if (savedNews.value.has(newsId)) {
    savedNews.value.delete(newsId);
  } else {
    savedNews.value.add(newsId);
  }
  savedNews.value = new Set(savedNews.value);
};

const isSaved = (newsId) => savedNews.value.has(newsId);

/* ========== UTILITIES ========== */
const formatDate = (dateStr) => {
  const d = new Date(dateStr);
  const monthNames = [
    "มกราคม", "กุมภาพันธ์", "มีนาคม", "เมษายน", "พฤษภาคม", "มิถุนายน",
    "กรกฎาคม", "สิงหาคม", "กันยายน", "ตุลาคม", "พฤศจิกายน", "ธันวาคม",
  ];
  return `${d.getDate()} ${monthNames[d.getMonth()]} ${d.getFullYear() + 543}`;
};
</script>

<style scoped>
/* ========== GLOBAL ========== */
@import url("https://fonts.googleapis.com/css2?family=Kanit:wght@400;700&display=swap");

* {
  box-sizing: border-box;
  font-family: "Kanit", "Prompt", "Sarabun", "Noto Sans Thai", sans-serif;
}

/* ========== SECTION ========== */
.news-events {
  padding: 20px 0;
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

/* ========== SEARCH & SORT ========== */
.search-sort {
  display: flex;
  align-items: center;
  margin: 25px 0;
  flex-wrap: wrap;
  gap: 12px;
}

.search-box {
  flex: 1;
  width: max-content;
}

.search-box input {
  width: 100%;
  padding: 10px 14px;
  border: 1px solid #eef0ef;
  border-radius: 12px;
  outline: none;
  font-size: 14px;
  background: #fff;
  transition: box-shadow .12s, border-color .12s;
}

.search-box input:focus {
  border-color: #037266;
  box-shadow: 0 8px 22px rgba(3, 114, 102, 0.06);
}

.btn-search {
  background: #037266;
  color: #fff;
  border: none;
  padding: 9px 18px;
  border-radius: 10px;
  cursor: pointer;
  font-weight: 600;
  box-shadow: 0 6px 18px rgba(3, 114, 102, 0.08);
}

.btn-search:hover {
  transform: translateY(-2px);
}

.btn-sort {
  background: #f1f1f1;
  border: none;
  padding: 8px 10px;
  border-radius: 10px;
  cursor: pointer;
}

/* ========== NEWS LIST & CARD ========== */
.news-list {
  display: flex;
  flex-direction: column;
  gap: 35px;
  max-height: var(--news-max-height);
  overflow-y: auto;
  padding-right: 10px;
  scrollbar-width: thin;
  scrollbar-color: #797979 #f1f1f1;
}

.news-list::-webkit-scrollbar {
  width: 8px;
}

.news-list::-webkit-scrollbar-track {
  background: #f1f1f1;
  border-radius: 8px;
}

.news-list::-webkit-scrollbar-thumb {
  background-color: #037266;
  border-radius: 8px;
}

.news-card {
  position: relative;
  display: flex;
  gap: 20px;
  padding: 20px;
  border-bottom: 1px solid #eee;
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

/* ========== BOOKMARK BUTTON (on card) ========== */
.bookmark-btn {
  position: absolute;
  top: 18px;
  right: 18px;
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
  0%, 100% {
    transform: scale(1);
  }
  50% {
    transform: scale(1.15);
  }
}

/* IMAGE LAYOUT */
.news-img img {
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

/* CONTENT */
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

/* ========== ACTION BUTTON (in side panel) ========== */
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
  0%, 100% {
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

/* Responsive */
@media (max-width: 900px) {
  .news-card {
    padding: 16px;
  }

  .side-panel-inner {
    padding: 28px 20px 21px 20px;
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

  .news-card {
    padding: 14px;
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