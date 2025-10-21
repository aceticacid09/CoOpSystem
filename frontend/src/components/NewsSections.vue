<!--  frontend/src/components/NewsSections.vue  -->
<template>
  <section class="news-events container">
    <!-- ใช้ TabNavigation แทน -->
    <TabNavigation v-model="activeTab" :tabs="tabs" />

    <SearchBar v-model="searchText" placeholder="ค้นหาหัวข้อข่าว" :is-ascending="isAscending"
      @toggle-sort="toggleSort" />

    <div v-if="isLoading" class="loading-state">
      <p>กำลังโหลดข่าวสาร...</p>
    </div>

    <div v-else-if="!filteredNews.length" class="no-results-state">
      <p>ไม่พบข่าวสารในหมวดหมู่นี้</p>
    </div>

    <div v-else class="news-list">
      <!-- ใช้ NewsCard แทน -->
      <NewsCard v-for="item in filteredNews" :key="item.id" :news="item" :show-bookmark="showBookmark"
        :is-bookmarked="isBookmarked(item.id)" @click="openModal" @toggle-bookmark="toggleBookmark" />
    </div>

    <!-- ใช้ NewsSidePanel แทน -->
    <NewsSidePanel :news="selectedNews" :show-actions="showBookmark" :is-saved="isSaved(selectedNews?.id)"
      @close="closeModal" @toggle-save="toggleSaveNews" />
  </section>
</template>

<script setup>
import { ref, computed, onMounted } from "vue";
import SearchBar from "../components/SearchBar.vue";
import TabNavigation from "../components/TabNavigation.vue";
import NewsCard from "../components/NewsCard.vue";
import NewsSidePanel from "../components/NewSidePanel.vue";
// import { getPublishedNews } from "../data/newsData.js";
import { API_CONFIG } from '../../config/api'; // <- relative path to frontend/config/api.js

// props
const props = defineProps({
  items: { type: Array, default: null },
  maxHeight: { type: String, default: null },
  showBookmark: { type: Boolean, default: false },
  initialBookmarked: { type: Array, default: () => [] },
  initialSaved: { type: Array, default: () => [] }
});

// tabs/state
const tabs = ["ข่าวทั้งหมด", "ประชาสัมพันธ์", "ข่าวด่วน", "กิจกรรม", "ประกาศผลการคัดเลือก"];
const activeTab = ref(tabs[0]);
const searchText = ref("");
const isAscending = ref(true);
const selectedNews = ref(null);
const bookmarkedNews = ref(new Set(props.initialBookmarked));
const savedNews = ref(new Set(props.initialSaved));

// local fetched items + loading/error
const itemsState = ref(null);
const isLoading = ref(false);
const fetchError = ref(null);

const transformAnnouncement = (announcement) => {
  return {
    id: announcement.post_id,
    title: announcement.title,
    description: announcement.description,
    date: announcement.created_at,
    category: "ประชาสัมพันธ์", // You might want to map this based on some criteria
    status: announcement.status,
    teacher: announcement.teacher,
    // Map attachments to images array if they exist
    images: announcement.attachments?.map(att => `/uploads/news/${att.filename}`) || []
  };
};

// Modify the onMounted function
onMounted(async () => {
  if (props.items) return;
  
  isLoading.value = true;
  try {
    console.log('Fetching from:', `${API_CONFIG.baseURL}${API_CONFIG.endpoints.announcements}`);
    const res = await fetch(`${API_CONFIG.baseURL}${API_CONFIG.endpoints.announcements}`);
    if (!res.ok) throw new Error(`Failed to fetch announcements: ${res.status}`);
    
    const data = await res.json();
    console.log('API Response:', data);
    
    // Transform the announcements data to match expected format
    const transformedAnnouncements = data.announcements.map(transformAnnouncement);
    itemsState.value = transformedAnnouncements;
    
    console.log('Processed items:', itemsState.value);
  } catch (err) {
    console.error("Failed to load announcements:", err);
    fetchError.value = err;
  } finally {
    isLoading.value = false;
  }
});

const filteredNews = computed(() => {
  const newsSource = props.items || itemsState.value;
  console.log('News source:', newsSource); // Debug source data
  
  const filtered = (newsSource || []).filter(
    n => {
      const matchesTab = activeTab.value === "ข่าวทั้งหมด" || n.category === activeTab.value;
      const matchesSearch = n.title?.includes(searchText.value) || n.description?.includes(searchText.value);
      return matchesTab && matchesSearch;
    }
  );
  
  console.log('Filtered results:', filtered); // Debug filtered results
  return isAscending.value ? filtered : filtered.slice().reverse();
});

const toggleSort = () => (isAscending.value = !isAscending.value);
const openModal = (item) => { selectedNews.value = { ...item }; };
const closeModal = () => (selectedNews.value = null);

const toggleBookmark = (newsId) => {
  if (bookmarkedNews.value.has(newsId)) bookmarkedNews.value.delete(newsId);
  else bookmarkedNews.value.add(newsId);
  bookmarkedNews.value = new Set(bookmarkedNews.value);
};

const isBookmarked = (newsId) => bookmarkedNews.value.has(newsId);

const toggleSaveNews = (newsId) => {
  if (savedNews.value.has(newsId)) savedNews.value.delete(newsId);
  else savedNews.value.add(newsId);
  savedNews.value = new Set(savedNews.value);
};

const isSaved = (newsId) => savedNews.value.has(newsId);
</script>


<style scoped>
@import url("https://fonts.googleapis.com/css2?family=Kanit:wght@400;700&display=swap");

* {
  box-sizing: border-box;
  font-family: "Kanit", "Prompt", "Sarabun", "Noto Sans Thai", sans-serif;
}

.news-events {
  padding: 20px 0;
}

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

.loading-state,
.no-results-state {
  text-align: center;
  padding: 50px 0;
  font-size: 16px;
  color: #555;
}
</style>