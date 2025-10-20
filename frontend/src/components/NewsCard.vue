<!-- frontend/src/components/NewsCard.vue -->
<template>
  <div class="news-card" @click="$emit('click', news)">
    <!-- Bookmark Button -->
    <button v-if="showBookmark" class="bookmark-btn" :class="{ bookmarked: isBookmarked }"
      @click="handleBookmarkClick" :aria-label="isBookmarked ? 'ยกเลิกบันทึก' : 'บันทึกข่าว'">
      <svg xmlns="http://www.w3.org/2000/svg" width="22" height="22" viewBox="0 0 24 24"
        :fill="isBookmarked ? '#ffd43b' : 'none'" stroke="currentColor" stroke-width="2" stroke-linecap="round"
        stroke-linejoin="round">
        <path d="M19 21l-7-5-7 5V5a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2z"></path>
      </svg>
    </button>

    <!-- IMAGE -->
    <div class="news-img" v-if="news.images?.length">
      <img v-if="news.images.length === 1" :src="news.images[0]" alt="news" />
      <div v-else class="image-grid" :class="`images-${Math.min(news.images.length, 4)}`">
        <div v-for="(img, i) in news.images.slice(0, 4)" :key="i" class="image-cell">
          <img :src="img" alt="news" />
          <div v-if="i === 3 && news.images.length > 4" class="image-overlay">
            +{{ news.images.length - 4 }}
          </div>
        </div>
      </div>
    </div>

    <!-- CONTENT -->
    <div class="news-content">
      <h3 class="news-title">{{ news.title }}</h3>
      <p class="news-date">
        <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" width="14" height="14" fill="none"
          stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
          <rect x="3" y="4" width="18" height="18" rx="2" ry="2" />
          <line x1="16" y1="2" x2="16" y2="6" />
          <line x1="8" y1="2" x2="8" y2="6" />
          <line x1="3" y1="10" x2="21" y2="10" />
        </svg>
        ประกาศเมื่อ {{ formatDate(news.date) }}
      </p>
      <p class="news-desc">{{ news.description }}</p>
    </div>
  </div>
</template>

<script setup>
const props = defineProps({
  news: {
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
  emit('toggle-bookmark', props.news.id);
};

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
  0%, 100% { transform: scale(1); }
  50% { transform: scale(1.15); }
}

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

.image-grid.images-3 .image-cell:nth-child(1) { grid-area: a; }
.image-grid.images-3 .image-cell:nth-child(2) { grid-area: b; }
.image-grid.images-3 .image-cell:nth-child(3) { grid-area: c; }

.image-grid.images-4 {
  grid-template-columns: repeat(2, 1fr);
  grid-template-rows: repeat(2, 80px);
  width: 220px;
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
}
</style>