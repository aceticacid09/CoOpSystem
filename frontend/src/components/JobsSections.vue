<!-- frontend/src/components/JobsSections.vue -->
<template>
  <section class="jobs-section container">
    <!-- Filters เหมือนเดิม -->
    <div class="jobs-filters">
      <div class="filters">
        <CustomDropdown label="ปีการศึกษา" v-model="selectedYear" :options="academicYears" />
        <CustomDropdown label="ภาควิชา" v-model="selectedDepartment" :options="departments" />
        <SegmentedControl label="ภาคที่" v-model="selectedSemester" :options="semesters" />
      </div>
      <SearchBar v-model="searchText" placeholder="ค้นหาชื่อตำแหน่ง, บริษัท หรือ ภาควิชา" :is-ascending="isAscending"
        @toggle-sort="toggleSort" />
    </div>

    <!-- Jobs List - ใช้ JobCard -->
    <div class="jobs-list">
      <JobCard v-for="job in filteredJobs" :key="job.id" :job="job" :show-bookmark="showBookmark"
        :is-bookmarked="isBookmarked(job.id)" @click="openPanel" @toggle-bookmark="toggleBookmark" />
    </div>

    <!-- ใช้ JobSidePanel -->
    <JobSidePanel :job="selectedJob" :show-actions="showActions" :is-applied="isApplied(selectedJob?.id)"
      :is-saved="isSaved(selectedJob?.id)" company-type-text="ดูงานทั้งหมด" @close="closePanel" @apply="applyJob"
      @toggle-save="toggleSaveJob" />
  </section>
</template>

<script setup>
import { ref, computed, onMounted, onBeforeUnmount } from "vue";
import jobsList from "../data/jobsData.js";
import { useRouter } from 'vue-router';
import SearchBar from "../components/SearchBar.vue";
import CustomDropdown from "../components/CustomDropdown.vue";
import JobCard from "../components/JobCard.vue";
import JobSidePanel from "../components/JobSidePanel.vue";
import SegmentedControl from "../components/SegmentedControl.vue";

const router = useRouter();

const props = defineProps({
  userRole: { type: String, default: null },
  showBookmark: { type: Boolean, default: false },
  showActions: { type: Boolean, default: false },
  initialBookmarked: { type: Array, default: () => [] },
  initialSaved: { type: Array, default: () => [] },
  initialApplied: { type: Array, default: () => [] }
});

// state เหมือนเดิม
const academicYears = [2568, 2567, 2566, 2565, 2564, 2563, 2562, 2561, 2560, 2559];
const departments = ["ทั้งหมด", "ภาควิชาคณิตศาสตร์", "ภาควิชาเคมี", "ภาควิชาชีววิทยา", "ภาควิชาฟิสิกส์", "ภาควิชาวิทยาศาสตร์สิ่งแวดล้อม", "ภาควิชาสถิติ", "ภาควิชาคอมพิวเตอร์", "ภาควิชาจุลชีววิทยา"];
const semesters = [1, 2, 3];

const selectedYear = ref(2568);
const selectedDepartment = ref("ทั้งหมด");
const selectedSemester = ref(1);
const searchText = ref("");
const isAscending = ref(true);
const selectedJob = ref(null);

const bookmarkedJobs = ref(new Set(props.initialBookmarked));
const appliedJobs = ref(new Set(props.initialApplied));
const savedJobs = ref(new Set(props.initialSaved));

const toggleSort = () => {
  isAscending.value = !isAscending.value;
};

const openPanel = (job) => {
  selectedJob.value = job;
};

const closePanel = () => {
  selectedJob.value = null;
};

const filteredJobs = computed(() => {
  let filtered = jobsList.filter(job => {
    if (selectedYear.value !== job.academicYear) return false;
    if (selectedDepartment.value !== "ทั้งหมด" && job.department !== selectedDepartment.value) return false;
    if (selectedSemester.value !== job.semester) return false;

    const txt = searchText.value.trim().toLowerCase();
    if (txt) {
      const searchFields = [job.title, job.company, job.department].join(" ").toLowerCase();
      if (!searchFields.includes(txt)) return false;
    }
    return true;
  });

  return isAscending.value ? filtered : [...filtered].reverse();
});

const toggleBookmark = (jobId) => {
  if (bookmarkedJobs.value.has(jobId)) {
    bookmarkedJobs.value.delete(jobId);
  } else {
    bookmarkedJobs.value.add(jobId);
  }
  bookmarkedJobs.value = new Set(bookmarkedJobs.value);
};

const isBookmarked = (jobId) => {
  return bookmarkedJobs.value.has(jobId);
};

const applyJob = (jobId) => {
  router.push({
    name: 'StudentJobApplication',
    params: { jobId: jobId }
  });
};

const toggleSaveJob = (jobId) => {
  if (savedJobs.value.has(jobId)) {
    savedJobs.value.delete(jobId);
  } else {
    savedJobs.value.add(jobId);
  }
  savedJobs.value = new Set(savedJobs.value);
};

const isApplied = (jobId) => appliedJobs.value.has(jobId);
const isSaved = (jobId) => savedJobs.value.has(jobId);
</script>

<style scoped>
@import url("https://fonts.googleapis.com/css2?family=Kanit:wght@400;600;700&display=swap");

.jobs-section {
  padding: 0;
}

* {
  box-sizing: border-box;
  font-family: "Kanit", "Prompt", "Sarabun", "Noto Sans Thai", sans-serif;
}

.jobs-filters {
  padding: 10px 0;
}

.filters {
  display: flex;
  gap: 24px;
  flex-wrap: wrap;
  align-items: center;
}

.jobs-list {
  margin-top: 10px;
  display: flex;
  flex-direction: column;
  gap: 32px;
}
</style>