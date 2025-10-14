<!-- frontend/src/components/JobsSections.vue -->
<template>
  <section class="jobs-section container">
    <!-- ========== FILTERS ========== -->
    <div class="jobs-filters">
      <div class="filters">
        <!-- Academic Year Dropdown -->
        <div class="filter-group">
          <label class="filter-label">ปีการศึกษา</label>
          <div class="custom-dropdown" ref="dropdownRef" @click="isDropdownOpen = !isDropdownOpen">
            <div class="dropdown-selected">
              <span>{{ selectedYear }}</span>
              <svg class="dropdown-arrow" :class="{ open: isDropdownOpen }" xmlns="http://www.w3.org/2000/svg"
                width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="#9CA3AF" stroke-width="2"
                stroke-linecap="round" stroke-linejoin="round">
                <polyline points="6 9 12 15 18 9"></polyline>
              </svg>
            </div>
            <transition name="dropdown-fade">
              <ul v-if="isDropdownOpen" class="dropdown-menu">
                <li v-for="year in academicYears" :key="year" @click.stop="selectYear(year)">
                  {{ year }}
                </li>
              </ul>
            </transition>
          </div>
        </div>
        <!-- Department Dropdown -->
        <div class="filter-group">
          <label class="filter-label">ภาควิชา</label>
          <div class="custom-dropdown" ref="deptDropdownRef" @click="isDeptDropdownOpen = !isDeptDropdownOpen">
            <div class="dropdown-selected">
              <span>{{ selectedDepartment }}</span>
              <svg class="dropdown-arrow" :class="{ open: isDeptDropdownOpen }" xmlns="http://www.w3.org/2000/svg"
                width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="#9CA3AF" stroke-width="2"
                stroke-linecap="round" stroke-linejoin="round">
                <polyline points="6 9 12 15 18 9"></polyline>
              </svg>
            </div>
            <transition name="dropdown-fade">
              <ul v-if="isDeptDropdownOpen" class="dropdown-menu">
                <li v-for="dept in departments" :key="dept" @click.stop="selectDepartment(dept)">
                  {{ dept }}
                </li>
              </ul>
            </transition>
          </div>
        </div>
        <!-- Semester Filter -->
        <div class="filter-group">
          <label class="filter-label">ภาคที่</label>
          <div class="segmented-control">
            <button v-for="sem in semesters" :key="sem" :class="{ active: selectedSemester === sem }"
              @click="selectedSemester = sem">
              {{ sem }}
            </button>
          </div>
        </div>
      </div>
    </div>

    <!-- ========== SEARCH & SORT ========== -->
    <div class="search-sort">
      <div class="search-box">
        <input type="text" v-model="searchText" placeholder="ค้นหาชื่อตำแหน่ง, บริษัท หรือภาควิชา" />
      </div>
      <button class="btn-search">ค้นหา</button>
      <button class="btn-sort">
        <svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" fill="none" stroke="#9CA3AF" stroke-width="2"
          stroke-linecap="round" stroke-linejoin="round">
          <path d="M3 6h18M3 12h12M3 18h6" />
        </svg>
      </button>
    </div>

    <!-- ========== JOB CARDS ========== -->
    <div class="jobs-list">
      <div v-for="job in filteredJobs" :key="job.id" class="job-card" @click="openPanel(job)">
        <div class="job-header">

          <!-- Bookmark Button -->
          <!-- Bookmark Button - แสดงเฉพาะเมื่อ showBookmark เป็น true -->
          <button v-if="showBookmark" class="bookmark-btn" :class="{ bookmarked: isBookmarked(job.id) }"
            @click="toggleBookmark(job.id, $event)" :aria-label="isBookmarked(job.id) ? 'ยกเลิกบันทึก' : 'บันทึกงาน'">
            <svg xmlns="http://www.w3.org/2000/svg" width="22" height="22" viewBox="0 0 24 24"
              :fill="isBookmarked(job.id) ? '#ffd43b' : 'none'" stroke="currentColor" stroke-width="2"
              stroke-linecap="round" stroke-linejoin="round">
              <path d="M19 21l-7-5-7 5V5a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2z"></path>
            </svg>
          </button>
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
                <span>ปิดรับสมัคร {{ job.closeDate }} (ระยะเวลา {{ job.applyDuration }})</span>
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

    <!-- ========== JOB DETAIL SIDE PANEL ========== -->
    <div v-if="selectedJob" class="side-panel-overlay" @click.self="closePanel">
      <div class="side-panel">
        <button class="side-close" @click="closePanel">&times;</button>
        <div class="side-panel-inner">
          <div class="side-header-top">
            <img class="side-logo" :src="selectedJob.logo" alt="logo" />
            <div class="side-title-block">
              <div class="side-title">{{ selectedJob.title }}</div>
              <div class="side-company-row">
                <span class="side-company-name">{{ selectedJob.company }}</span>
                <span class="side-company-sep"></span>
                <span class="side-company-all-jobs">ดูงานทั้งหมด</span>
              </div>
              <span class="side-department-label">
                {{ selectedJob.department }}
              </span>
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
              ปิดรับสมัคร {{ selectedJob.closeDate }} (ระยะเวลา
              {{ selectedJob.applyDuration }})
            </span>
          </div>
          <div class="side-meta-row">
            <span class="side-meta-item">
              <svg width="22" height="22" fill="none" stroke="#9CA3AF" stroke-width="2" stroke-linecap="round"
                stroke-linejoin="round" viewBox="0 0 24 24">
                <path d="M22 10L12 4 2 10l10 6 10-6z"></path>
                <path d="M6 12v5c0 1.1 2.7 2 6 2s6-.9 6-2v-5"></path>
              </svg>
              ปีการศึกษา {{ selectedJob.academicYear }} ภาคที่
              {{ selectedJob.semester }}
            </span>
          </div>
          <!-- Action Buttons -->
          <!-- Action Buttons - แสดงเฉพาะเมื่อ showActions เป็น true (นักศึกษาเท่านั้น) -->
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

            <button class="btn-save" :class="{ saved: isSaved(selectedJob.id) }" @click="toggleSaveJob(selectedJob.id)">
              <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24"
                :fill="isSaved(selectedJob.id) ? 'currentColor' : 'none'" stroke="currentColor" stroke-width="2"
                stroke-linecap="round" stroke-linejoin="round">
                <path d="M19 21l-7-5-7 5V5a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2z"></path>
              </svg>
              <span>{{ isSaved(selectedJob.id) ? 'บันทึกแล้ว' : 'บันทึก' }}</span>
            </button>
          </div>
          <!-- Images Preview -->
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
  import { ref, onMounted, onBeforeUnmount, computed, watch } from "vue";
  import jobsList from "../data/jobsData.js";
  import { useRouter } from 'vue-router';

const router = useRouter();

const props = defineProps({
  userRole: {
    type: String,
    default: null
  },
  showBookmark: {
    type: Boolean,
    default: false
  },
  showActions: {
    type: Boolean,
    default: false
  },
  // เพิ่ม props ใหม่เหล่านี้
  initialBookmarked: {
    type: Array,
    default: () => [] // เช่น [1, 3, 5] = งาน id 1,3,5 ถูกบันทึกไว้แล้ว
  },
  initialSaved: {
    type: Array,
    default: () => []
  },
  initialApplied: {
    type: Array,
    default: () => [] // เช่น [2, 4] = งาน id 2,4 สมัครไว้แล้ว
  }
});

const academicYears = [2568, 2567, 2566, 2565, 2564, 2563, 2562, 2561, 2560, 2559];
const departments = ["ทั้งหมด", "ภาควิชาคณิตศาสตร์", "ภาควิชาเคมี", "ภาควิชาชีววิทยา", "ภาควิชาฟิสิกส์", "ภาควิชาวิทยาศาสตร์สิ่งแวดล้อม", "ภาควิชาสถิติ", "ภาควิชาคอมพิวเตอร์", "ภาควิชาจุลชีววิทยา"];
const semesters = [1, 2, 3];

const selectedYear = ref(2568);
const isDropdownOpen = ref(false);
const selectYear = (year) => { selectedYear.value = year; isDropdownOpen.value = false; };

const selectedDepartment = ref("ทั้งหมด");
const isDeptDropdownOpen = ref(false);
const selectDepartment = (dept) => { selectedDepartment.value = dept; isDeptDropdownOpen.value = false; };

const selectedSemester = ref(1);
const searchText = ref("");
const dropdownRef = ref(null);
const deptDropdownRef = ref(null);
const selectedJob = ref(null);

const selectedImageIndex = ref(0);

watch(selectedJob, (job) => {
  selectedImageIndex.value = 0;
});

const openPanel = (job) => {
  selectedJob.value = job;
};
const closePanel = () => {
  selectedJob.value = null;
};

const handleClickOutside = (event) => {
  if (dropdownRef.value && !dropdownRef.value.contains(event.target)) isDropdownOpen.value = false;
  if (deptDropdownRef.value && !deptDropdownRef.value.contains(event.target)) isDeptDropdownOpen.value = false;
};

onMounted(() => document.addEventListener("click", handleClickOutside));
onBeforeUnmount(() => document.removeEventListener("click", handleClickOutside));

function truncateText(text, len) {
  if (!text) return "";
  return text.length <= len ? text : text.slice(0, len);
}

const filteredJobs = computed(() => {
  return jobsList.filter(job => {
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
});

const bookmarkedJobs = ref(new Set(props.initialBookmarked));
const appliedJobs = ref(new Set(props.initialApplied));
const savedJobs = ref(new Set(props.initialSaved));

const toggleBookmark = (jobId, event) => {
  event.stopPropagation(); // Prevent opening panel
  if (bookmarkedJobs.value.has(jobId)) {
    bookmarkedJobs.value.delete(jobId);
  } else {
    bookmarkedJobs.value.add(jobId);
  }
  // Trigger reactivity
  bookmarkedJobs.value = new Set(bookmarkedJobs.value);
};

const isBookmarked = (jobId) => {
  return bookmarkedJobs.value.has(jobId);
};

const applyJob = (jobId) => {
  // นำทางไปหน้าสมัครงาน
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

.filter-group {
  display: flex;
  flex-direction: column;
  gap: 6px;
}

.filter-label {
  font-size: 13px;
  color: #6b6b6b;
  font-weight: 500;
}

.custom-dropdown {
  position: relative;
  width: 250px;
  cursor: pointer;
  user-select: none;
}

.dropdown-selected {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 9px 14px;
  background: #fff;
  border: 1px solid #e5e7e6;
  border-radius: 10px;
  font-size: 14px;
  color: #333;
  font-weight: 500;
  transition: all 0.2s ease;
  box-shadow: 0 1px 3px rgba(10, 10, 10, 0.02);
}

.dropdown-selected:hover {
  border-color: #037266;
  box-shadow: 0 4px 12px rgba(3, 114, 102, 0.08);
}

.dropdown-arrow {
  transition: transform 0.25s ease;
  flex-shrink: 0;
  margin-left: 8px;
}

.dropdown-arrow.open {
  transform: rotate(180deg);
}

.dropdown-menu {
  position: absolute;
  top: calc(100% + 6px);
  left: 0;
  width: 100%;
  background: #fff;
  border: 1px solid #e5e7e6;
  border-radius: 10px;
  box-shadow: 0 8px 24px rgba(6, 20, 18, 0.12);
  max-height: 240px;
  overflow-y: auto;
  z-index: 100;
  list-style: none;
  margin: 0;
  padding: 6px;
}

.dropdown-menu li {
  padding: 9px 12px;
  font-size: 14px;
  color: #333;
  cursor: pointer;
  border-radius: 6px;
  transition: all 0.15s ease;
  font-weight: 500;
  white-space: nowrap;
}

.dropdown-menu li:hover {
  background: #f0f8f7;
  color: #037266;
}

.dropdown-fade-enter-active,
.dropdown-fade-leave-active {
  transition: opacity 0.2s ease, transform 0.2s ease;
  transform-origin: top;
}

.dropdown-fade-enter-from {
  opacity: 0;
  transform: translateY(-8px) scale(0.95);
}

.dropdown-fade-leave-to {
  opacity: 0;
  transform: translateY(-4px) scale(0.98);
}

.segmented-control {
  display: inline-flex;
  border-radius: 12px;
  background: #f7f9f8;
  padding: 4px;
  gap: 4px;
  border: 1px solid #eef4f2;
  box-shadow: 0 1px 3px rgba(10, 10, 10, 0.02);
}

.segmented-control button {
  min-width: 44px;
  padding: 8px 12px;
  border-radius: 8px;
  border: none;
  background: transparent;
  cursor: pointer;
  font-weight: 600;
  color: #556;
  transition: background .18s, color .18s, transform .08s;
}

.segmented-control button.active {
  background: linear-gradient(180deg, #037266, #026b5b);
  color: #fff;
  box-shadow: 0 6px 16px rgba(3, 114, 102, 0.12);
  transform: translateY(-1px);
}

.search-sort {
  display: flex;
  align-items: center;
  gap: 12px;
  margin-bottom: 20px;
  flex-wrap: wrap;
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
  border-color: #0b6b57;
  box-shadow: 0 8px 22px rgba(11, 107, 87, 0.06);
}

.btn-search {
  background: #0b6b57;
  color: #fff;
  border: none;
  padding: 9px 18px;
  border-radius: 10px;
  cursor: pointer;
  font-weight: 600;
  box-shadow: 0 6px 18px rgba(11, 107, 87, 0.08);
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

.jobs-list {
  margin-top: 10px;
  display: flex;
  flex-direction: column;
  gap: 32px;
}

.job-card {
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
  min-height: 100vh;
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

.side-company-all-jobs {
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
  background: none;
  padding: 0;
  border-radius: 0;
  white-space: pre-line;
}

/* Responsive */
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
}

/* ========== BOOKMARK BUTTON (on card) ========== */
.job-card {
  position: relative;
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

  0%,
  100% {
    transform: scale(1);
  }

  50% {
    transform: scale(1.15);
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

.btn-apply:active:not(:disabled) {
  transform: translateY(0);
  box-shadow: 0 2px 8px rgba(3, 114, 102, 0.2);
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

  0%,
  100% {
    transform: scale(1);
  }

  50% {
    transform: scale(1.05);
  }
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

/* Icon animations */
.btn-apply svg,
.btn-save svg {
  transition: transform 0.3s ease;
}

.btn-apply:hover:not(:disabled) svg,
.btn-save:hover svg {
  transform: scale(1.1);
}

.btn-apply.applied svg {
  animation: checkmarkBounce 0.6s ease;
}

@keyframes checkmarkBounce {

  0%,
  100% {
    transform: scale(1);
  }

  30% {
    transform: scale(1.3);
  }

  60% {
    transform: scale(0.9);
  }
}

/* Responsive adjustments */
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

  .side-action-buttons {
    flex-direction: column;
  }

  .btn-apply,
  .btn-save {
    min-width: 100%;
  }
}
</style>