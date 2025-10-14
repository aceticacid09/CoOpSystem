<!-- frontend/src/pages/Student/StudentHistoryJobs.vue -->
<template>
  <DashboardLayout role="student">
    <section class="history-section container">
      <!-- ========== FILTERS & SEARCH ========== -->
      <div class="history-controls">
        <div class="filter-row">
          <!-- Company Dropdown -->
          <div class="filter-group">
            <label class="filter-label">บริษัท</label>
            <div class="custom-dropdown" ref="companyDropdownRef"
              @click="isCompanyDropdownOpen = !isCompanyDropdownOpen">
              <div class="dropdown-selected">
                <span>{{ selectedCompany }}</span>
                <svg class="dropdown-arrow" :class="{ open: isCompanyDropdownOpen }" xmlns="http://www.w3.org/2000/svg"
                  width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="#9CA3AF" stroke-width="2"
                  stroke-linecap="round" stroke-linejoin="round">
                  <polyline points="6 9 12 15 18 9"></polyline>
                </svg>
              </div>
              <transition name="dropdown-fade">
                <ul v-if="isCompanyDropdownOpen" class="dropdown-menu">
                  <li v-for="company in companies" :key="company" @click.stop="selectCompany(company)">
                    {{ company }}
                  </li>
                </ul>
              </transition>
            </div>
          </div>

          <!-- Search Box -->
          <div class="search-box">
            <input type="text" v-model="searchText" placeholder="ค้นหาชื่อตำแหน่ง หรือบริษัท" />
          </div>

          <!-- Search Button -->
          <button class="btn-search">ค้นหา</button>

          <!-- Filter Button -->
          <button class="btn-sort">
            <svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" fill="none" stroke="#9CA3AF" stroke-width="2"
              stroke-linecap="round" stroke-linejoin="round">
              <path d="M3 6h18M3 12h12M3 18h6" />
            </svg>
          </button>
        </div>
      </div>

      <!-- ========== HISTORY CARDS ========== -->
      <div class="history-list">
        <div v-for="record in filteredRecords" :key="record.id" class="history-card">
          <div class="card-header">
            <div class="header-info">
              <h3 class="job-title">{{ record.jobTitle }}</h3>
              <p class="company-name">{{ record.company }}</p>
              <p class="application-date">สมัครเมื่อ {{ formatDate(record.applicationDate) }}</p>
            </div>
            <div class="status-badge" :class="getStatusClass(record.stages)">
              {{ getCurrentStage(record.stages) }}
            </div>
          </div>

          <div class="timeline-container">
            <div class="timeline-wrapper">
              <div v-for="(stage, index) in record.stages" :key="index" class="timeline-item">
                <div class="timeline-dot" :class="stage.status">
                  <svg v-if="stage.status === 'passed'" xmlns="http://www.w3.org/2000/svg" width="18" height="18"
                    viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="3" stroke-linecap="round"
                    stroke-linejoin="round">
                    <polyline points="20 6 9 17 4 12"></polyline>
                  </svg>
                  <span v-else-if="stage.status === 'current'">?</span>
                  <svg v-else-if="stage.status === 'rejected'" xmlns="http://www.w3.org/2000/svg" width="18" height="18"
                    viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round"
                    stroke-linejoin="round">
                    <line x1="18" y1="6" x2="6" y2="18"></line>
                    <line x1="6" y1="6" x2="18" y2="18"></line>
                  </svg>
                </div>
                <div class="timeline-label">{{ stage.name }}</div>
                <div v-if="index < record.stages.length - 1" class="timeline-connector"
                  :class="getConnectorClass(record.stages, index)">
                </div>
              </div>
            </div>
          </div>

          <div class="card-actions">
            <button class="btn-detail">ดูประกาศงาน</button>
            <button class="btn-application">ดูใบสมัครงาน</button>
          </div>
        </div>
      </div>
    </section>
  </DashboardLayout>
</template>

<script setup>
import { ref, computed, onMounted, onBeforeUnmount } from "vue";
import DashboardLayout from '../../components/DashboardLayout.vue';

const isCompanyDropdownOpen = ref(false);
const selectedCompany = ref("ทั้งหมด");
const searchText = ref("");
const companyDropdownRef = ref(null);

const monthsThai = [
  "มกราคม",
  "กุมภาพันธ์",
  "มีนาคม",
  "เมษายน",
  "พฤษภาคม",
  "มิถุนายน",
  "กรกฎาคม",
  "สิงหาคม",
  "กันยายน",
  "ตุลาคม",
  "พฤศจิกายน",
  "ธันวาคม"
];

const applicationRecords = ref([
  {
    id: 1,
    jobTitle: "Frontend Developer (In-tern)",
    company: "Google (Thailand) Company Limited",
    applicationDate: "11/09/68",
    stages: [
      { name: "กดสมัคร", status: "passed" },
      { name: "ตรวจสอบเอกสาร", status: "current" },
      { name: "สัมภาษณ์", status: "pending" },
      { name: "รอพิจารณา", status: "pending" },
      { name: "ประกาศผล", status: "pending" }
    ]
  },
  {
    id: 2,
    jobTitle: "Frontend Developer",
    company: "AAA ltd.",
    applicationDate: "11/09/68",
    stages: [
      { name: "กดสมัคร", status: "passed" },
      { name: "ตรวจสอบเอกสาร", status: "passed" },
      { name: "สัมภาษณ์", status: "current" },
      { name: "รอพิจารณา", status: "pending" },
      { name: "ประกาศผล", status: "pending" }
    ]
  },
  {
    id: 3,
    jobTitle: "Frontent Developer",
    company: "ABCD ltd.",
    applicationDate: "11/09/68",
    stages: [
      { name: "กดสมัคร", status: "passed" },
      { name: "ตรวจสอบเอกสาร", status: "passed" },
      { name: "สัมภาษณ์", status: "passed" },
      { name: "รอพิจารณา", status: "current" },
      { name: "ประกาศผล", status: "pending" }
    ]
  },
  {
    id: 4,
    jobTitle: "UI/UX Designer",
    company: "Tech Innovate Co.",
    applicationDate: "10/09/68",
    stages: [
      { name: "กดสมัคร", status: "passed" },
      { name: "ตรวจสอบเอกสาร", status: "rejected" },
      { name: "สัมภาษณ์", status: "pending" },
      { name: "รอพิจารณา", status: "pending" },
      { name: "ประกาศผล", status: "pending" }
    ]
  },
  {
    id: 5,
    jobTitle: "Backend Developer",
    company: "Digital Solutions Ltd.",
    applicationDate: "09/09/68",
    stages: [
      { name: "กดสมัคร", status: "passed" },
      { name: "ตรวจสอบเอกสาร", status: "passed" },
      { name: "สัมภาษณ์", status: "passed" },
      { name: "รอพิจารณา", status: "passed" },
      { name: "ประกาศผล", status: "passed" }
    ]
  },
  {
    id: 6,
    jobTitle: "Data Analyst",
    company: "Analytics Pro Group",
    applicationDate: "08/09/68",
    stages: [
      { name: "กดสมัคร", status: "passed" },
      { name: "ตรวจสอบเอกสาร", status: "passed" },
      { name: "สัมภาษณ์", status: "current" },
      { name: "รอพิจารณา", status: "pending" },
      { name: "ประกาศผล", status: "pending" }
    ]
  }
]);

const companies = computed(() => {
  const uniqueCompanies = new Set(applicationRecords.value.map(r => r.company));
  return ["ทั้งหมด", ...Array.from(uniqueCompanies)];
});

const selectCompany = (company) => {
  selectedCompany.value = company;
  isCompanyDropdownOpen.value = false;
};

const handleClickOutside = (event) => {
  if (companyDropdownRef.value && !companyDropdownRef.value.contains(event.target)) {
    isCompanyDropdownOpen.value = false;
  }
};

const formatDate = (dateStr) => {
  const parts = dateStr.split("/");
  const day = parts[0];
  const month = parseInt(parts[1]) - 1;
  const year = parseInt(parts[2]) + 2500;

  return `${day} ${monthsThai[month]} ${year}`;
};

const getCurrentStage = (stages) => {
  const currentStage = stages.find(s => s.status === "current");
  const rejectedStage = stages.find(s => s.status === "rejected");

  if (rejectedStage) {
    return `ไม่ผ่าน${rejectedStage.name}`;
  }

  if (currentStage) {
    return `${currentStage.name}`;
  }

  const passedStages = stages.filter(s => s.status === "passed");
  if (passedStages.length === stages.length) {
    return "ผ่านการคัดเลือก";
  }

  return "รอดำเนินการ";
};

const getStatusClass = (stages) => {
  const rejectedStage = stages.find(s => s.status === "rejected");
  if (rejectedStage) return "rejected";

  const currentStage = stages.find(s => s.status === "current");
  const passedStages = stages.filter(s => s.status === "passed");

  if (passedStages.length === stages.length) return "completed";

  if (currentStage) {
    const currentIndex = stages.findIndex(s => s.status === "current");
    if (currentIndex === 0 || currentIndex === 1) return "applying";
    if (currentIndex === 2) return "interview";
    if (currentIndex === 3) return "considering";
  }

  return "applying";
};

const getConnectorClass = (stages, index) => {
  const currentStage = stages[index];
  const nextStage = stages[index + 1];

  if (currentStage.status === "passed" && (nextStage.status === "passed" || nextStage.status === "current")) {
    return "passed";
  }

  if (currentStage.status === "passed" && nextStage.status === "rejected") {
    return "rejected";
  }

  return "pending";
};

const filteredRecords = computed(() => {
  return applicationRecords.value.filter(record => {
    if (selectedCompany.value !== "ทั้งหมด" && record.company !== selectedCompany.value) return false;
    if (searchText.value.trim()) {
      const txt = searchText.value.toLowerCase();
      return record.jobTitle.toLowerCase().includes(txt) || record.company.toLowerCase().includes(txt);
    }
    return true;
  });
});

onMounted(() => {
  document.addEventListener("click", handleClickOutside);
});

onBeforeUnmount(() => {
  document.removeEventListener("click", handleClickOutside);
});
</script>

<style scoped>
@import url("https://fonts.googleapis.com/css2?family=Kanit:wght@400;500;600;700&display=swap");

.history-section {
  padding: 0;
}

* {
  box-sizing: border-box;
  font-family: "Kanit", "Prompt", "Sarabun", "Noto Sans Thai", sans-serif;
}

.history-controls {
  padding: 10px 0;
  margin-bottom: 20px;
}

.filter-row {
  display: flex;
  gap: 12px;
  align-items: flex-end;
  flex-wrap: wrap;
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
  transition: box-shadow 0.12s, border-color 0.12s;
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
  font-size: 14px;
  box-shadow: 0 6px 18px rgba(11, 107, 87, 0.08);
  transition: transform 0.2s;
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
  transition: background 0.2s, transform 0.2s;
  display: flex;
  align-items: center;
  justify-content: center;
}

.btn-sort:hover {
  background: #e8e8e8;
  transform: translateY(-2px);
}

/* ========== HISTORY CARDS ========== */
.history-list {
  display: flex;
  flex-direction: column;
  gap: 18px;
  margin-top: 20px;
}

.history-card {
  background: #fff;
  border-radius: 16px;
  border: 1px solid #f0f0f0;
  padding: 26px 30px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.04);
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
}

.history-card:hover {
  box-shadow: 0 8px 24px rgba(3, 114, 102, 0.12);
  transform: translateY(-2px);
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: flex-start;
  gap: 20px;
  margin-bottom: 28px;
}

.header-info {
  flex: 1;
}

.job-title {
  font-size: 19px;
  font-weight: 700;
  color: #1a1a1a;
  margin: 0 0 8px 0;
  letter-spacing: -0.2px;
  line-height: 1.3;
}

.company-name {
  font-size: 14px;
  color: #666;
  font-weight: 500;
  margin: 0 0 6px 0;
  letter-spacing: 0.1px;
}

.application-date {
  font-size: 13px;
  color: #999;
  margin: 0;
  font-weight: 400;
}

/* ========== STATUS BADGES (ปรับสีใหม่) ========== */
.status-badge {
  padding: 8px 20px;
  border-radius: 20px;
  font-size: 13px;
  font-weight: 600;
  white-space: nowrap;
  transition: all 0.3s ease;
  flex-shrink: 0;
}

/* ตรวจสอบเอกสาร - น้ำเงินอ่อน */
.status-badge.applying {
  background: #dbeafe;
  color: #1e40af;
}

/* นัดสัมภาษณ์ - ฟ้าอ่อน */
.status-badge.interview {
  background: #e0f2fe;
  color: #0369a1;
}

/* รอพิจารณา - เหลืองอ่อน (เหมือนเดิม) */
.status-badge.considering {
  background: #fef3c7;
  color: #92400e;
}

/* ผ่านการคัดเลือก - เขียว */
.status-badge.completed {
  background: #d1fae5;
  color: #065f46;
}

/* ไม่ผ่านการคัดเลือก - แดง */
.status-badge.rejected {
  background: #fee2e2;
  color: #991b1b;
}

/* ========== TIMELINE ========== */
.timeline-container {
  position: relative;
  margin-bottom: 28px;
  padding: 0 10px;
}

.timeline-wrapper {
  display: flex;
  justify-content: space-between;
  align-items: flex-start;
  gap: 6px;
  position: relative;
  z-index: 2;
}

.timeline-item {
  display: flex;
  flex-direction: column;
  align-items: center;
  flex: 1;
  position: relative;
}

.timeline-dot {
  width: 40px;
  height: 40px;
  border-radius: 50%;
  background: #f8f8f8;
  border: 3px solid #e8e8e8;
  display: flex;
  align-items: center;
  justify-content: center;
  margin-bottom: 14px;
  transition: all 0.3s ease;
  flex-shrink: 0;
  position: relative;
  z-index: 3;
  font-weight: 700;
  font-size: 20px;
  box-shadow: 0 2px 6px rgba(0, 0, 0, 0.08);
}

.timeline-dot.passed {
  background: linear-gradient(135deg, #10b981 0%, #059669 100%);
  border-color: #10b981;
  color: #fff;
  box-shadow: 0 3px 12px rgba(16, 185, 129, 0.3);
}

.timeline-dot.current {
  background: linear-gradient(135deg, #fbbf24 0%, #f59e0b 100%);
  border-color: #f59e0b;
  color: #fff;
  box-shadow: 0 0 0 4px rgba(245, 158, 11, 0.2), 0 3px 12px rgba(245, 158, 11, 0.3);
  animation: pulseIndicator 2s infinite;
  font-size: 22px;
}

.timeline-dot.rejected {
  background: linear-gradient(135deg, #ef4444 0%, #dc2626 100%);
  border-color: #ef4444;
  color: #fff;
  box-shadow: 0 3px 12px rgba(239, 68, 68, 0.3);
}

.timeline-dot.pending {
  background: #f8f8f8;
  border-color: #e8e8e8;
  color: #ccc;
  box-shadow: 0 2px 6px rgba(0, 0, 0, 0.05);
}

@keyframes pulseIndicator {

  0%,
  100% {
    box-shadow: 0 0 0 4px rgba(245, 158, 11, 0.2), 0 3px 12px rgba(245, 158, 11, 0.3);
    transform: scale(1);
  }

  50% {
    box-shadow: 0 0 0 8px rgba(245, 158, 11, 0.1), 0 3px 12px rgba(245, 158, 11, 0.3);
    transform: scale(1.02);
  }
}

.timeline-label {
  font-size: 13px;
  color: #555;
  font-weight: 500;
  text-align: center;
  word-break: break-word;
  line-height: 1.4;
  letter-spacing: 0.2px;
  max-width: 90px;
}

.timeline-connector {
  position: absolute;
  top: 19px;
  left: calc(50% + 22px);
  right: calc(-50% + 22px);
  height: 3px;
  background: #e8e8e8;
  z-index: 1;
  transition: all 0.3s ease;
}

.timeline-connector.passed {
  background: linear-gradient(to right, #10b981 0%, #10b981 100%);
  box-shadow: 0 1px 4px rgba(16, 185, 129, 0.2);
}

.timeline-connector.rejected {
  background: linear-gradient(to right, #10b981 0%, #ef4444 100%);
}

.timeline-item:last-child .timeline-connector {
  display: none;
}

/* ========== ACTION BUTTONS ========== */
.card-actions {
  display: flex;
  gap: 12px;
}

.btn-detail,
.btn-application {
  flex: 1;
  padding: 12px 22px;
  border: none;
  border-radius: 12px;
  font-size: 14px;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.25s ease;
  letter-spacing: 0.3px;
}

.btn-detail {
  background: linear-gradient(135deg, #f0f9f7 0%, #e8f5f2 100%);
  color: #037266;
  border: 2px solid #c5e4e0;
}

.btn-detail:hover {
  background: linear-gradient(135deg, #e0f2f0 0%, #d8eae7 100%);
  border-color: #037266;
  transform: translateY(-2px);
  box-shadow: 0 6px 16px rgba(3, 114, 102, 0.2);
}

.btn-detail:active {
  transform: translateY(0);
}

.btn-application {
  background: #fff;
  color: #666;
  border: 2px solid #e0e0e0;
}

.btn-application:hover {
  background: #f9f9f9;
  color: #333;
  border-color: #999;
  transform: translateY(-2px);
  box-shadow: 0 6px 16px rgba(0, 0, 0, 0.1);
}

.btn-application:active {
  transform: translateY(0);
}

/* Responsive */
@media (max-width: 900px) {
  .filter-row {
    gap: 10px;
  }

  .custom-dropdown {
    width: 200px;
  }

  .search-box {
    min-width: 250px;
    max-width: 400px;
  }

  .history-card {
    padding: 20px 24px;
  }

  .card-header {
    flex-direction: column;
    gap: 14px;
  }

  .status-badge {
    align-self: flex-start;
  }
}

@media (max-width: 650px) {
  .filter-row {
    flex-direction: column;
    gap: 12px;
  }

  .filter-group,
  .search-box {
    width: 100%;
  }

  .custom-dropdown {
    width: 100%;
  }

  .search-box {
    max-width: 100%;
  }

  .btn-search,
  .btn-sort {
    width: 100%;
  }

  .filter-row {
    display: grid;
    grid-template-columns: 1fr 1fr;
    gap: 12px;
  }

  .filter-group {
    grid-column: span 2;
  }

  .search-box {
    grid-column: 1;
  }

  .btn-search {
    grid-column: 2;
  }

  .btn-sort {
    grid-column: span 2;
  }

  .history-card {
    padding: 16px 18px;
  }

  .job-title {
    font-size: 16px;
  }

  .status-badge {
    padding: 8px 18px;
    font-size: 13px;
  }

  .timeline-container {
    padding: 0 5px;
  }

  .timeline-wrapper {
    gap: 3px;
  }

  .timeline-dot {
    width: 36px;
    height: 36px;
  }

  .timeline-dot svg {
    width: 16px;
    height: 16px;
  }

  .timeline-dot.current {
    font-size: 18px;
  }

  .timeline-label {
    font-size: 11px;
    max-width: 70px;
  }

  .timeline-connector {
    top: 17px;
    height: 2px;
    left: calc(50% + 20px);
    right: calc(-50% + 20px);
  }

  .card-actions {
    flex-direction: column;
    gap: 10px;
  }

  .btn-detail,
  .btn-application {
    width: 100%;
  }
}
</style>
