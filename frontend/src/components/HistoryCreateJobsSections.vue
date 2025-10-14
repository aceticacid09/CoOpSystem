<!-- frontend/src/components/CompanySections/HistoryCreateJobsSections.vue -->
<template>
  <section class="history-jobs-section container">
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

        <!-- Status Dropdown -->
        <div class="filter-group">
          <label class="filter-label">สถานะ</label>
          <div class="custom-dropdown" ref="statusDropdownRef" @click="isStatusDropdownOpen = !isStatusDropdownOpen">
            <div class="dropdown-selected">
              <span>{{ selectedStatus }}</span>
              <svg class="dropdown-arrow" :class="{ open: isStatusDropdownOpen }" xmlns="http://www.w3.org/2000/svg"
                width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="#9CA3AF" stroke-width="2"
                stroke-linecap="round" stroke-linejoin="round">
                <polyline points="6 9 12 15 18 9"></polyline>
              </svg>
            </div>
            <transition name="dropdown-fade">
              <ul v-if="isStatusDropdownOpen" class="dropdown-menu">
                <li v-for="status in statuses" :key="status" @click.stop="selectStatus(status)">
                  {{ status }}
                </li>
              </ul>
            </transition>
          </div>
        </div>
      </div>
    </div>

    <!-- ========== SEARCH & SORT ========== -->
    <div class="search-sort">
      <div class="search-box">
        <svg class="search-icon" xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="#9CA3AF" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
          <circle cx="11" cy="11" r="8"></circle>
          <path d="m21 21-4.35-4.35"></path>
        </svg>
        <input type="text" v-model="searchText" placeholder="ค้นหาชื่อตำแหน่งงาน" />
      </div>
      <button class="btn-search" @click="handleSearch">ค้นหา</button>
      <button class="btn-sort">
        <svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" fill="none" stroke="#9CA3AF" stroke-width="2"
          stroke-linecap="round" stroke-linejoin="round">
          <path d="M3 6h18M3 12h12M3 18h6" />
        </svg>
      </button>
    </div>

    <!-- ========== TABLE ========== -->
    <div class="table-wrapper">
      <div class="table-container">
        <table class="jobs-table">
          <thead>
            <tr>
              <th>ชื่อตำแหน่งงาน</th>
              <th>ภาควิชา</th>
              <th>สถานะ</th>
              <th>ระยะการรับสมัคร</th>
              <th>วันที่ลงประกาศ</th>
              <th>ผู้สมัครงาน</th>
              <th>การดำเนินการ</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="job in paginatedJobs" :key="job.id">
              <td class="job-title">{{ job.title }}</td>
              <td class="department">{{ job.department }}</td>
              <td>
                <span class="status-badge" :class="getStatusClass(job.status)">
                  {{ job.status }}
                </span>
              </td>
              <td class="duration">{{ job.duration }} วัน</td>
              <td class="date">{{ job.postedDate }}</td>
              <td class="applicants">{{ job.applicants }} คน</td>
              <td>
                <button class="btn-view" @click="viewJob(job.id)">
                  ดูประกาศงาน
                </button>
              </td>
            </tr>
          </tbody>
        </table>
      </div>

      <!-- Empty State -->
      <div v-if="filteredJobs.length === 0" class="empty-state">
        <svg xmlns="http://www.w3.org/2000/svg" width="64" height="64" viewBox="0 0 24 24" fill="none" stroke="#D1D5DB" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round">
          <path d="M20 21v-2a4 4 0 0 0-4-4H8a4 4 0 0 0-4 4v2"></path>
          <circle cx="12" cy="7" r="4"></circle>
        </svg>
        <p>ไม่พบรายการประกาศงาน</p>
        <span>ลองเปลี่ยนตัวกรองหรือคำค้นหาของคุณ</span>
      </div>

      <!-- Pagination Controls -->
      <div v-if="filteredJobs.length > 0" class="pagination-wrapper">
        <div class="rows-per-page">
          <span>แสดง</span>
          <select v-model="rowsPerPage" @change="currentPage = 1">
            <option :value="10">10</option>
            <option :value="25">25</option>
            <option :value="50">50</option>
            <option :value="100">100</option>
          </select>
          <span>รายการต่อหน้า</span>
        </div>

        <div class="pagination-info">
          <span>{{ startIndex + 1 }} - {{ endIndex }} จาก {{ filteredJobs.length }} รายการ</span>
        </div>

        <div class="pagination-controls">
          <button 
            class="pagination-btn" 
            :disabled="currentPage === 1"
            @click="currentPage--"
          >
            <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
              <polyline points="15 18 9 12 15 6"></polyline>
            </svg>
          </button>

          <div class="page-numbers">
            <button 
              v-for="page in visiblePages" 
              :key="page"
              class="page-btn"
              :class="{ active: page === currentPage }"
              @click="currentPage = page"
            >
              {{ page }}
            </button>
          </div>

          <button 
            class="pagination-btn" 
            :disabled="currentPage === totalPages"
            @click="currentPage++"
          >
            <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
              <polyline points="9 18 15 12 9 6"></polyline>
            </svg>
          </button>
        </div>
      </div>
    </div>
  </section>
</template>

<script setup>
import { ref, computed, onMounted, onBeforeUnmount } from "vue";

// Data
const academicYears = [2568, 2567, 2566, 2565, 2564, 2563, 2562, 2561, 2560, 2559];
const departments = ["ทั้งหมด", "ภาควิชาคณิตศาสตร์", "ภาควิชาเคมี", "ภาควิชาชีววิทยา", "ภาควิชาฟิสิกส์", "ภาควิชาวิทยาศาสตร์สิ่งแวดล้อม", "ภาควิชาสถิติ", "ภาควิชาคอมพิวเตอร์", "ภาควิชาจุลชีววิทยา"];
const semesters = [1, 2, 3];
const statuses = ["ทั้งหมด", "รอตรวจสอบ", "ผ่านการตรวจสอบ", "ไม่ผ่านการตรวจสอบ"];

// Mock data for jobs
const jobs = ref([
  { id: 1, year: 2568, semester: 1, title: "Software Engineer", department: "ภาควิชาคอมพิวเตอร์", status: "ผ่านการตรวจสอบ", duration: 30, postedDate: "15/01/2568", applicants: 24 },
  { id: 2, year: 2568, semester: 2, title: "Data Scientist", department: "ภาควิชาสถิติ", status: "รอตรวจสอบ", duration: 15, postedDate: "20/01/2568", applicants: 12 },
  { id: 3, year: 2568, semester: 1, title: "Research Assistant", department: "ภาควิชาชีววิทยา", status: "ผ่านการตรวจสอบ", duration: 45, postedDate: "10/01/2568", applicants: 8 },
  { id: 4, year: 2567, semester: 2, title: "Lab Technician", department: "ภาควิชาเคมี", status: "ไม่ผ่านการตรวจสอบ", duration: 7, postedDate: "18/01/2568", applicants: 0 },
  { id: 5, year: 2567, semester: 2, title: "Physics Instructor", department: "ภาควิชาฟิสิกส์", status: "ผ่านการตรวจสอบ", duration: 60, postedDate: "05/01/2568", applicants: 35 },
  { id: 6, year: 2568, semester: 3, title: "Environmental Analyst", department: "ภาควิชาวิทยาศาสตร์สิ่งแวดล้อม", status: "รอตรวจสอบ", duration: 20, postedDate: "22/01/2568", applicants: 5 },
  { id: 7, year: 25687, semester: 1, title: "Mathematics Tutor", department: "ภาควิชาคณิตศาสตร์", status: "ผ่านการตรวจสอบ", duration: 30, postedDate: "12/01/2568", applicants: 18 },
  { id: 8, year: 2568, semester: 1, title: "Microbiologist", department: "ภาควิชาจุลชีววิทยา", status: "ผ่านการตรวจสอบ", duration: 25, postedDate: "08/01/2568", applicants: 21 },
  { id: 9, year: 2568, semester: 2, title: "Quality Assurance", department: "ภาควิชาเคมี", status: "รอตรวจสอบ", duration: 14, postedDate: "25/01/2568", applicants: 7 },
  { id: 10, year: 2568, semester: 1, title: "IT Support", department: "ภาควิชาคอมพิวเตอร์", status: "ผ่านการตรวจสอบ", duration: 30, postedDate: "16/01/2568", applicants: 42 },
  { id: 11, year: 2568, semester: 3, title: "Statistician", department: "ภาควิชาสถิติ", status: "ไม่ผ่านการตรวจสอบ", duration: 10, postedDate: "28/01/2568", applicants: 3 },
  { id: 12, year: 2568, semester: 3, title: "Biology Teacher", department: "ภาควิชาชีววิทยา", status: "ผ่านการตรวจสอบ", duration: 40, postedDate: "03/01/2568", applicants: 29 },
  { id: 13, year: 2568, semester: 1, title: "Chemistry Instructor", department: "ภาควิชาเคมี", status: "รอตรวจสอบ", duration: 18, postedDate: "21/01/2568", applicants: 11 },
]);

// State for filters
const selectedYear = ref(2568);
const isDropdownOpen = ref(false);
const selectedDepartment = ref("ทั้งหมด");
const isDeptDropdownOpen = ref(false);
const selectedSemester = ref(1);
const selectedStatus = ref("ทั้งหมด");
const isStatusDropdownOpen = ref(false);
const searchText = ref("");

// Pagination state
const currentPage = ref(1);
const rowsPerPage = ref(10);

// Refs
const dropdownRef = ref(null);
const deptDropdownRef = ref(null);
const statusDropdownRef = ref(null);

// Computed
const filteredJobs = computed(() => {
  return jobs.value.filter(job => {
    // เพิ่มเงื่อนไขปีการศึกษาและภาคการศึกษา
    const matchesYear = job.year === selectedYear.value;
    const matchesSemester = job.semester === selectedSemester.value;

    const matchesDept = selectedDepartment.value === "ทั้งหมด" || job.department === selectedDepartment.value;
    const matchesStatus = selectedStatus.value === "ทั้งหมด" || job.status === selectedStatus.value;
    const matchesSearch = !searchText.value || job.title.toLowerCase().includes(searchText.value.toLowerCase())

    // ต้อง return ทุกเงื่อนไข
    return matchesYear && matchesSemester && matchesDept && matchesStatus && matchesSearch;
  });
});

const totalPages = computed(() => Math.ceil(filteredJobs.value.length / rowsPerPage.value));

const startIndex = computed(() => (currentPage.value - 1) * rowsPerPage.value);
const endIndex = computed(() => Math.min(startIndex.value + rowsPerPage.value, filteredJobs.value.length));

const paginatedJobs = computed(() => {
  return filteredJobs.value.slice(startIndex.value, endIndex.value);
});

const visiblePages = computed(() => {
  const pages = [];
  const maxVisible = 5;
  let start = Math.max(1, currentPage.value - Math.floor(maxVisible / 2));
  let end = Math.min(totalPages.value, start + maxVisible - 1);
  
  if (end - start + 1 < maxVisible) {
    start = Math.max(1, end - maxVisible + 1);
  }
  
  for (let i = start; i <= end; i++) {
    pages.push(i);
  }
  
  return pages;
});

// Methods
const selectYear = (year) => {
  selectedYear.value = year;
  isDropdownOpen.value = false;
};

const selectDepartment = (dept) => {
  selectedDepartment.value = dept;
  isDeptDropdownOpen.value = false;
  currentPage.value = 1;
};

const selectStatus = (status) => {
  selectedStatus.value = status;
  isStatusDropdownOpen.value = false;
  currentPage.value = 1;
};

const handleSearch = () => {
  currentPage.value = 1;
};

const viewJob = (id) => {
  console.log("View job:", id);
  // Navigate to job details page
};

const getStatusClass = (status) => {
  const statusMap = {
    "รอตรวจสอบ": "pending",
    "ผ่านการตรวจสอบ": "approved",
    "ไม่ผ่านการตรวจสอบ": "rejected"
  };
  return statusMap[status] || "";
};

// Click outside handler
const handleClickOutside = (event) => {
  if (dropdownRef.value && !dropdownRef.value.contains(event.target)) {
    isDropdownOpen.value = false;
  }
  if (deptDropdownRef.value && !deptDropdownRef.value.contains(event.target)) {
    isDeptDropdownOpen.value = false;
  }
  if (statusDropdownRef.value && !statusDropdownRef.value.contains(event.target)) {
    isStatusDropdownOpen.value = false;
  }
};

onMounted(() => {
  document.addEventListener("click", handleClickOutside);
});

onBeforeUnmount(() => {
  document.removeEventListener("click", handleClickOutside);
});
</script>

<style scoped>
@import url("https://fonts.googleapis.com/css2?family=Kanit:wght@400;500;600;700&display=swap");

.history-jobs-section {
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
  margin-bottom: 24px;
  flex-wrap: wrap;
}

.search-box {
  flex: 1;
  position: relative;
}

.search-icon {
  position: absolute;
  left: 14px;
  top: 50%;
  transform: translateY(-50%);
  pointer-events: none;
}

.search-box input {
  width: 100%;
  padding: 10px 14px 10px 40px;
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
  padding: 10px 24px;
  border-radius: 10px;
  cursor: pointer;
  font-weight: 600;
  font-size: 14px;
  box-shadow: 0 6px 18px rgba(11, 107, 87, 0.08);
  transition: all 0.2s ease;
}

.btn-search:hover {
  background: #0a5f4d;
  transform: translateY(-2px);
  box-shadow: 0 8px 22px rgba(11, 107, 87, 0.15);
}

.btn-sort {
  background: #f1f1f1;
  border: none;
  padding: 8px 10px;
  border-radius: 10px;
  cursor: pointer;
}

/* Table Styles */
.table-wrapper {
  background: #fff;
  border-radius: 14px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.04);
  overflow: hidden;
  border: 1px solid #f0f0f0;
}

.table-container {
  overflow-x: auto;
}

.jobs-table {
  width: 100%;
  border-collapse: collapse;
  min-width: 900px;
}

.jobs-table thead {
  background: linear-gradient(180deg, #f8fafa 0%, #f2f6f5 100%);
  border-bottom: 2px solid #e5e7e6;
}

.jobs-table th {
  padding: 16px 20px;
  text-align: left;
  font-size: 13px;
  font-weight: 700;
  color: #4a5568;
  text-transform: uppercase;
  letter-spacing: 0.5px;
}

.jobs-table tbody tr {
  border-bottom: 1px solid #f0f0f0;
  transition: all 0.2s ease;
}

.jobs-table tbody tr:hover {
  background: #fafbfb;
  box-shadow: 0 2px 8px rgba(3, 114, 102, 0.04);
}

.jobs-table tbody tr:last-child {
  border-bottom: none;
}

.jobs-table td {
  padding: 18px 20px;
  font-size: 14px;
  color: #2d3748;
  vertical-align: middle;
}

.job-title {
  font-weight: 600;
  color: #1a202c;
}

.department {
  color: #4a5568;
  font-size: 13px;
}

.status-badge {
  display: inline-block;
  padding: 6px 14px;
  border-radius: 20px;
  font-size: 12px;
  font-weight: 600;
  text-align: center;
  white-space: nowrap;
}

.status-badge.pending {
  background: #fef3c7;
  color: #92400e;
}

.status-badge.approved {
  background: #d1fae5;
  color: #065f46;
}

.status-badge.rejected {
  background: #fee2e2;
  color: #991b1b;
}

.duration,
.date,
.applicants {
  color: #4a5568;
  font-weight: 500;
}

.btn-view {
  display: inline-flex;
  align-items: center;
  gap: 6px;
  padding: 8px 16px;
  background: #f0f8f7;
  color: #037266;
  border: 1px solid #d1e9e6;
  border-radius: 8px;
  cursor: pointer;
  font-size: 13px;
  font-weight: 600;
  transition: all 0.2s ease;
}

.btn-view:hover {
  background: #037266;
  color: #fff;
  border-color: #037266;
  transform: translateY(-1px);
  box-shadow: 0 4px 12px rgba(3, 114, 102, 0.15);
}

/* Empty State */
.empty-state {
  padding: 80px 20px;
  text-align: center;
}

.empty-state svg {
  margin-bottom: 20px;
}

.empty-state p {
  font-size: 18px;
  font-weight: 600;
  color: #4a5568;
  margin-bottom: 8px;
}

.empty-state span {
  font-size: 14px;
  color: #9ca3af;
}

/* Pagination */
.pagination-wrapper {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 20px 24px;
  background: #fafbfb;
  border-top: 1px solid #f0f0f0;
  flex-wrap: wrap;
  gap: 16px;
}

.rows-per-page {
  display: flex;
  align-items: center;
  gap: 10px;
  font-size: 14px;
  color: #4a5568;
}

.rows-per-page select {
  padding: 6px 32px 6px 12px;
  border: 1px solid #e5e7e6;
  border-radius: 8px;
  font-size: 14px;
  font-weight: 600;
  color: #2d3748;
  background: #fff;
  cursor: pointer;
  transition: all 0.2s ease;
  appearance: none;
  background-image: url("data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' width='12' height='12' viewBox='0 0 24 24' fill='none' stroke='%239CA3AF' stroke-width='2' stroke-linecap='round' stroke-linejoin='round'%3E%3Cpolyline points='6 9 12 15 18 9'%3E%3C/polyline%3E%3C/svg%3E");
  background-repeat: no-repeat;
  background-position: right 10px center;
}

.rows-per-page select:hover {
  border-color: #037266;
  box-shadow: 0 2px 8px rgba(3, 114, 102, 0.08);
}

.rows-per-page select:focus {
  outline: none;
  border-color: #037266;
  box-shadow: 0 0 0 3px rgba(3, 114, 102, 0.1);
}

.pagination-info {
  font-size: 14px;
  color: #4a5568;
  font-weight: 500;
}

.pagination-controls {
  display: flex;
  align-items: center;
  gap: 8px;
}

.pagination-btn {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 36px;
  height: 36px;
  border: 1px solid #e5e7e6;
  border-radius: 8px;
  background: #fff;
  cursor: pointer;
  transition: all 0.2s ease;
}

.pagination-btn:hover:not(:disabled) {
  background: #f0f8f7;
  border-color: #037266;
  color: #037266;
}

.pagination-btn:disabled {
  opacity: 0.4;
  cursor: not-allowed;
}

.page-numbers {
  display: flex;
  gap: 4px;
}

.page-btn {
  min-width: 36px;
  height: 36px;
  padding: 0 10px;
  border: 1px solid #e5e7e6;
  border-radius: 8px;
  background: #fff;
  color: #4a5568;
  font-size: 14px;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.2s ease;
}

.page-btn:hover {
  background: #f0f8f7;
  border-color: #037266;
  color: #037266;
}

.page-btn.active {
  background: linear-gradient(180deg, #037266, #026b5b);
  border-color: #037266;
  color: #fff;
  box-shadow: 0 4px 12px rgba(3, 114, 102, 0.15);
}

/* Responsive */
@media (max-width: 900px) {
  .filters {
    gap: 16px;
  }

  .custom-dropdown {
    width: 220px;
  }

  .jobs-table {
    font-size: 13px;
  }

  .jobs-table th,
  .jobs-table td {
    padding: 14px 16px;
  }
}

@media (max-width: 768px) {
  .pagination-wrapper {
    flex-direction: column;
    align-items: stretch;
  }

  .rows-per-page,
  .pagination-info {
    justify-content: center;
  }

  .pagination-controls {
    justify-content: center;
  }
}

@media (max-width: 650px) {
  .filters {
    flex-direction: column;
    align-items: stretch;
  }

  .custom-dropdown {
    width: 100%;
  }

  .segmented-control {
    width: 100%;
    justify-content: space-between;
  }

  .search-sort {
    flex-direction: column;
  }

  .search-box {
    width: 100%;
  }

  .btn-search {
    width: 100%;
  }

  .jobs-table th,
  .jobs-table td {
    padding: 12px;
    font-size: 12px;
  }

  .status-badge {
    font-size: 11px;
    padding: 4px 10px;
  }

  .btn-view {
    font-size: 12px;
    padding: 6px 12px;
  }

  .page-numbers {
    flex-wrap: wrap;
  }
}
</style>