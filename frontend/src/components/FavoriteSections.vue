<template>
  <section class="favorites-section container">
    <!-- ใช้ TabNavigation -->
    <TabNavigation v-model="activeTab" :tabs="tabs" />

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
        <NewsCard
          v-for="item in savedNews"
          :key="item.id"
          :news="item"
          :show-bookmark="true"
          :is-bookmarked="true"
          @click="openNewsModal"
          @toggle-bookmark="removeNewsBookmark(item.id, $event)"
        />
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
        <JobCard
          v-for="job in savedJobs"
          :key="job.id"
          :job="job"
          :show-bookmark="true"
          :is-bookmarked="true"
          @click="openJobPanel"
          @toggle-bookmark="removeJobBookmark(job.id, $event)"
        />
      </div>
    </div>

    <!-- News Side Panel -->
    <NewsSidePanel
      :news="selectedNews"
      :show-actions="showActions"
      :is-saved="isSavedNews(selectedNews?.id)"
      @close="closeNewsModal"
      @toggle-save="toggleSaveNews"
    />

    <!-- Job Side Panel -->
    <JobSidePanel
      :job="selectedJob"
      :show-actions="showActions"
      :is-applied="isApplied(selectedJob?.id)"
      :is-saved="isSavedJob(selectedJob?.id)"
      company-type-text="ดูรายละเอียดงาน"
      @close="closeJobPanel"
      @apply="applyJob"
      @toggle-save="toggleSaveJob"
    />
  </section>
</template>

<script setup>
import { ref, computed } from "vue";
import { newsList } from "../data/newsData.js";
import jobsList from "../data/jobsData.js";
import TabNavigation from "../components/TabNavigation.vue";
import NewsCard from "../components/NewsCard.vue";
import JobCard from "../components/JobCard.vue";
import NewsSidePanel from "../components/NewSidePanel.vue";
import JobSidePanel from "../components/JobSidePanel.vue";

const props = defineProps({
  savedNewsIds: { type: Array, default: () => [] },
  savedJobsIds: { type: Array, default: () => [] },
  showBookmark: { type: Boolean, default: true },
  showActions: { type: Boolean, default: true }
});

const emit = defineEmits(['remove-news-bookmark', 'remove-job-bookmark']);

const tabs = ["ข่าวสารที่บันทึก", "งานที่บันทึก"];
const activeTab = ref(tabs[0]);

const selectedNews = ref(null);
const selectedJob = ref(null);

const appliedJobs = ref(new Set());
const savedNewsInPanel = ref(new Set(props.savedNewsIds));
const savedJobsInPanel = ref(new Set(props.savedJobsIds));

const savedNews = computed(() => {
  return newsList.filter(news => props.savedNewsIds.includes(news.id));
});

const savedJobs = computed(() => {
  return jobsList.filter(job => props.savedJobsIds.includes(job.id));
});

const openNewsModal = (item) => {
  selectedNews.value = { ...item };
};

const closeNewsModal = () => {
  selectedNews.value = null;
};

const openJobPanel = (job) => {
  selectedJob.value = job;
};

const closeJobPanel = () => {
  selectedJob.value = null;
};

const removeNewsBookmark = (newsId, event) => {
  emit('remove-news-bookmark', newsId);
};

const removeJobBookmark = (jobId, event) => {
  emit('remove-job-bookmark', jobId);
};

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

.favorites-content {
  padding: 20px 0;
}

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

.news-list {
  display: flex;
  flex-direction: column;
  gap: 35px;
}

.jobs-list {
  display: flex;
  flex-direction: column;
  gap: 32px;
}

</style>