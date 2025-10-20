<!-- frontend/src/components/CompanySections/JobApplicationsList.vue -->
<template>
    <div class="job-applications-page">
        <div class="page-header">
            <div class="header-content">
                <h1>จัดการใบสมัครงาน</h1>
                <p class="page-subtitle">ตรวจสอบ จัดการ และติดตามสถานะของผู้สมัครงานทั้งหมด</p>
            </div>
        </div>

        <div class="stats-overview">
            <div class="stat-card">
                <div class="stat-icon total">
                    <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none"
                        stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                        <path d="M16 21v-2a4 4 0 0 0-4-4H5a4 4 0 0 0-4 4v2"></path>
                        <circle cx="8.5" cy="7" r="4"></circle>
                        <polyline points="17 11 19 13 23 9"></polyline>
                    </svg>
                </div>
                <div class="stat-info">
                    <span class="stat-label">ทั้งหมด</span>
                    <span class="stat-value">{{ totalCount }}</span>
                </div>
            </div>

            <!-- <div class="stat-card">
                <div class="stat-icon reviewing">
                    <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none"
                        stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                        <path d="M14 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8z"></path>
                        <polyline points="14 2 14 8 20 8"></polyline>
                        <line x1="16" y1="13" x2="8" y2="13"></line>
                        <line x1="16" y1="17" x2="8" y2="17"></line>
                        <polyline points="10 9 9 9 8 9"></polyline>
                    </svg>
                </div>
                <div class="stat-info">
                    <span class="stat-label">ตรวจสอบเอกสาร</span>
                    <span class="stat-value">{{ reviewingCount }}</span>
                </div>
            </div> -->

            <div class="stat-card">
                <div class="stat-icon interview">
                    <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none"
                        stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                        <path d="M17 21v-2a4 4 0 0 0-4-4H5a4 4 0 0 0-4 4v2"></path>
                        <circle cx="9" cy="7" r="4"></circle>
                        <path d="M23 21v-2a4 4 0 0 0-3-3.87"></path>
                        <path d="M16 3.13a4 4 0 0 1 0 7.75"></path>
                    </svg>
                </div>
                <div class="stat-info">
                    <span class="stat-label">นัดสัมภาษณ์</span>
                    <span class="stat-value">{{ interviewCount }}</span>
                </div>
            </div>

            <div class="stat-card">
                <div class="stat-icon pending">
                    <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none"
                        stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                        <circle cx="12" cy="12" r="10"></circle>
                        <polyline points="12 6 12 12 16 14"></polyline>
                    </svg>
                </div>
                <div class="stat-info">
                    <span class="stat-label">รอพิจารณา</span>
                    <span class="stat-value">{{ pendingCount }}</span>
                </div>
            </div>

            <div class="stat-card">
                <div class="stat-icon approved">
                    <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none"
                        stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                        <path d="M22 11.08V12a10 10 0 1 1-5.93-9.14"></path>
                        <polyline points="22 4 12 14.01 9 11.01"></polyline>
                    </svg>
                </div>
                <div class="stat-info">
                    <span class="stat-label">ผ่านการคัดเลือก</span>
                    <span class="stat-value">{{ approvedCount }}</span>
                </div>
            </div>

            <!-- <div class="stat-card">
                <div class="stat-icon rejected">
                    <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none"
                        stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                        <circle cx="12" cy="12" r="10"></circle>
                        <line x1="15" y1="9" x2="9" y2="15"></line>
                        <line x1="9" y1="9" x2="15" y2="15"></line>
                    </svg>
                </div>
                <div class="stat-info">
                    <span class="stat-label">ไม่ผ่านการคัดเลือก</span>
                    <span class="stat-value">{{ rejectedCount }}</span>
                </div>
            </div> -->
        </div>

        <div class="filter-section">
            <div class="tabs">
                <button v-for="tab in tabs" :key="tab.value" class="tab-btn"
                    :class="{ active: activeTab === tab.value }" @click="activeTab = tab.value">
                    {{ tab.label }}
                </button>
            </div>

            <div class="filters-row">
                <div class="filter-group">
                    <label class="filter-label">ปีการศึกษา</label>
                    <div class="custom-dropdown" ref="yearDropdownRef"
                        @click="isYearDropdownOpen = !isYearDropdownOpen">
                        <div class="dropdown-selected">
                            <span>{{ selectedYear }}</span>
                            <svg class="dropdown-arrow" :class="{ open: isYearDropdownOpen }"
                                xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24"
                                fill="none" stroke="#9CA3AF" stroke-width="2" stroke-linecap="round"
                                stroke-linejoin="round">
                                <polyline points="6 9 12 15 18 9"></polyline>
                            </svg>
                        </div>
                        <transition name="dropdown-fade">
                            <ul v-if="isYearDropdownOpen" class="dropdown-menu">
                                <li v-for="year in academicYears" :key="year" @click.stop="selectYear(year)">
                                    {{ year }}
                                </li>
                            </ul>
                        </transition>
                    </div>
                </div>

                <div class="filter-group">
                    <label class="filter-label">ภาคที่</label>
                    <div class="segmented-control">
                        <button v-for="sem in semesters" :key="sem" :class="{ active: selectedSemester === sem }"
                            @click="selectedSemester = sem">
                            {{ sem }}
                        </button>
                    </div>
                </div>

                <div class="filter-group">
                    <label class="filter-label">ตำแหน่งงาน</label>
                    <div class="custom-dropdown" ref="jobDropdownRef" @click="isJobDropdownOpen = !isJobDropdownOpen">
                        <div class="dropdown-selected">
                            <span>{{ selectedJob }}</span>
                            <svg class="dropdown-arrow" :class="{ open: isJobDropdownOpen }"
                                xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24"
                                fill="none" stroke="#9CA3AF" stroke-width="2" stroke-linecap="round"
                                stroke-linejoin="round">
                                <polyline points="6 9 12 15 18 9"></polyline>
                            </svg>
                        </div>
                        <transition name="dropdown-fade">
                            <ul v-if="isJobDropdownOpen" class="dropdown-menu">
                                <li v-for="job in jobPositions" :key="job" @click.stop="selectJob(job)">
                                    {{ job }}
                                </li>
                            </ul>
                        </transition>
                    </div>
                </div>
            </div>

            <div class="search-sort">
                <div class="search-box">
                    <svg class="search-icon" xmlns="http://www.w3.org/2000/svg" width="18" height="18"
                        viewBox="0 0 24 24" fill="none" stroke="#9CA3AF" stroke-width="2" stroke-linecap="round"
                        stroke-linejoin="round">
                        <circle cx="11" cy="11" r="8"></circle>
                        <path d="m21 21-4.35-4.35"></path>
                    </svg>
                    <input type="text" v-model="searchText" placeholder="ค้นหาชื่อนักศึกษา หรือตำแหน่งงาน" />
                </div>
                <button class="btn-search" @click="handleSearch">ค้นหา</button>
                <button class="btn-sort" @click="toggleSort"
                    :title="sortAscending ? 'เรียงจากเก่าสุด' : 'เรียงจากใหม่สุด'">
                    <svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" fill="none" stroke="#9CA3AF"
                        stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                        <path d="M3 6h18M3 12h12M3 18h6" />
                    </svg>
                </button>
            </div>
        </div>

        <div v-if="isLoading" class="loading-state">
            <div class="spinner"></div>
            <p>กำลังโหลดข้อมูล...</p>
        </div>

        <div v-else-if="filteredApplications.length === 0" class="empty-state">
            <svg xmlns="http://www.w3.org/2000/svg" width="64" height="64" viewBox="0 0 24 24" fill="none"
                stroke="currentColor" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round">
                <path d="M20 21v-2a4 4 0 0 0-4-4H8a4 4 0 0 0-4 4v2"></path>
                <circle cx="12" cy="7" r="4"></circle>
            </svg>
            <h3>ไม่พบใบสมัครงาน</h3>
            <p>{{ getEmptyMessage() }}</p>
        </div>

        <div v-else class="table-wrapper">
            <div class="table-container">
                <table class="applications-table">
                    <thead>
                        <tr>
                            <th style="width: 50px;"></th>
                            <th>ชื่อผู้สมัคร</th>
                            <th>ตำแหน่งงาน</th>
                            <th>สถานะ</th>
                            <th>วันที่สมัคร</th>
                        </tr>
                    </thead>
                    <tbody>
                        <template v-for="application in paginatedApplications" :key="application.id">
                            <tr class="clickable-row" @click="toggleRow(application.id)"
                                :class="{ 'expanded': expandedRows.has(application.id) }">
                                <td class="expand-cell">
                                    <svg class="expand-icon" :class="{ 'rotated': expandedRows.has(application.id) }"
                                        xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24"
                                        fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round"
                                        stroke-linejoin="round">
                                        <polyline points="9 18 15 12 9 6"></polyline>
                                    </svg>
                                </td>
                                <td class="applicant-name">{{ application.studentName }}</td>
                                <td class="job-position">{{ application.jobTitle }}</td>
                                <td>
                                    <span class="status-badge" :class="getStatusClass(application.status)">
                                        <span class="status-dot"></span>
                                        {{ application.status }}
                                    </span>
                                </td>
                                <td class="date">{{ application.appliedDate }}</td>
                            </tr>

                            <tr v-if="expandedRows.has(application.id)" class="details-row"
                                :class="getStatusBackgroundClass(application.status)">
                                <td colspan="5">
                                    <div class="details-content">
                                        <div class="student-info-section">
                                            <div class="info-header">
                                                <h4>ข้อมูลผู้สมัครงาน</h4>
                                            </div>
                                            <div class="student-name">{{ application.studentName }}</div>
                                            <div class="contact-list">
                                                <div class="contact-item">
                                                    <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16"
                                                        viewBox="0 0 24 24" fill="none" stroke="currentColor"
                                                        stroke-width="2" stroke-linecap="round" stroke-linejoin="round"
                                                        class="contact-icon">
                                                        <rect x="2" y="4" width="20" height="16" rx="2"></rect>
                                                        <path d="m2 7 10 6 10-6"></path>
                                                    </svg>
                                                    <span>{{ application.email }}</span>
                                                </div>
                                                <div class="contact-item">
                                                    <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16"
                                                        viewBox="0 0 24 24" fill="none" stroke="currentColor"
                                                        stroke-width="2" stroke-linecap="round" stroke-linejoin="round"
                                                        class="contact-icon">
                                                        <path
                                                            d="M22 16.92v3a2 2 0 0 1-2.18 2 19.79 19.79 0 0 1-8.63-3.07 19.5 19.5 0 0 1-6-6 19.79 19.79 0 0 1-3.07-8.67A2 2 0 0 1 4.11 2h3a2 2 0 0 1 2 1.72 12.84 12.84 0 0 0 .7 2.81 2 2 0 0 1-.45 2.11L8.09 9.91a16 16 0 0 0 6 6l1.27-1.27a2 2 0 0 1 2.11-.45 12.84 12.84 0 0 0 2.81.7A2 2 0 0 1 22 16.92z">
                                                        </path>
                                                    </svg>
                                                    <span>{{ application.phone }}</span>
                                                </div>
                                            </div>
                                        </div>

                                        <div class="actions-section">
                                            <!-- สถานะ: ตรวจสอบเอกสาร -->
                                            <template v-if="application.status === 'ตรวจสอบเอกสาร'">
                                                <button class="btn-status btn-approve"
                                                    @click.stop="updateStatus(application.id, 'นัดสัมภาษณ์')">
                                                    นัดสัมภาษณ์
                                                </button>
                                                <button class="btn-status btn-reject"
                                                    @click.stop="updateStatus(application.id, 'ไม่ผ่านการคัดเลือก')">
                                                    ไม่ผ่านการคัดเลือก
                                                </button>
                                            </template>

                                            <!-- สถานะ: นัดสัมภาษณ์ -->
                                            <template v-else-if="application.status === 'นัดสัมภาษณ์'">
                                                <button class="btn-status btn-interview-done"
                                                    @click.stop="updateStatus(application.id, 'รอพิจารณา')">
                                                    สัมภาษณ์เสร็จสิ้น
                                                </button>
                                                <button class="btn-status btn-reject"
                                                    @click.stop="updateStatus(application.id, 'ไม่ผ่านการคัดเลือก')">
                                                    ไม่ผ่านการคัดเลือก
                                                </button>
                                            </template>

                                            <!-- สถานะ: รอพิจารณา -->
                                            <template v-else-if="application.status === 'รอพิจารณา'">
                                                <button class="btn-status btn-approve"
                                                    @click.stop="updateStatus(application.id, 'ผ่านการคัดเลือก')">
                                                    ผ่านการคัดเลือก
                                                </button>
                                                <button class="btn-status btn-reject"
                                                    @click.stop="updateStatus(application.id, 'ไม่ผ่านการคัดเลือก')">
                                                    ไม่ผ่านการคัดเลือก
                                                </button>
                                            </template>

                                            <!-- ปุ่มดูใบสมัครงาน (แสดงเสมอ) -->
                                            <button class="btn-status btn-view-application"
                                                @click.stop="viewApplication(application.id)">
                                                ดูใบสมัครงาน
                                            </button>
                                        </div>
                                    </div>
                                </td>
                            </tr>
                        </template>
                    </tbody>
                </table>
            </div>

            <div v-if="filteredApplications.length > itemsPerPage" class="pagination">
                <button class="btn-page" @click="currentPage = 1" :disabled="currentPage === 1">
                    <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none"
                        stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                        <polyline points="11 17 6 12 11 7"></polyline>
                        <polyline points="18 17 13 12 18 7"></polyline>
                    </svg>
                </button>
                <button class="btn-page" @click="currentPage--" :disabled="currentPage === 1">
                    <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none"
                        stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                        <polyline points="15 18 9 12 15 6"></polyline>
                    </svg>
                </button>

                <template v-for="page in visiblePages" :key="page">
                    <button v-if="page !== '...'" class="btn-page" :class="{ active: currentPage === page }"
                        @click="currentPage = page">
                        {{ page }}
                    </button>
                    <span v-else class="page-ellipsis">...</span>
                </template>

                <button class="btn-page" @click="currentPage++" :disabled="currentPage === totalPages">
                    <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none"
                        stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                        <polyline points="9 18 15 12 9 6"></polyline>
                    </svg>
                </button>
                <button class="btn-page" @click="currentPage = totalPages" :disabled="currentPage === totalPages">
                    <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none"
                        stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                        <polyline points="13 17 18 12 13 7"></polyline>
                        <polyline points="6 17 11 12 6 7"></polyline>
                    </svg>
                </button>
            </div>
        </div>
    </div>
</template>

<script setup>
import { ref, computed, watch, onMounted, onBeforeUnmount } from "vue";

// =======================================================
// STATE
// =======================================================
const isLoading = ref(false);
const activeTab = ref('all');
const searchText = ref('');
const selectedYear = ref(2568);
const selectedSemester = ref(1);
const selectedJob = ref('ทั้งหมด');
const sortAscending = ref(false);
const currentPage = ref(1);
const itemsPerPage = ref(10);

// Dropdown states
const isYearDropdownOpen = ref(false);
const isJobDropdownOpen = ref(false);
const yearDropdownRef = ref(null);
const jobDropdownRef = ref(null);

// Expanded rows state
const expandedRows = ref(new Set());

// Data
const academicYears = [2568, 2567, 2566, 2565, 2564, 2563, 2562, 2561, 2560, 2559];
const semesters = [1, 2, 3];
const jobPositions = [
    'ทั้งหมด',
    'Frontend Developer (In-tern)',
    'Backend Developer (In-tern)',
    'Mobile Developer',
    'Data Analytics',
    'UI/UX Designer'
];

// Mock data
const mockApplicationsData = ref([
    {
        id: 1,
        year: 2568,
        semester: 1,
        studentName: "สมชาย ใจดี",
        jobTitle: "Frontend Developer (In-tern)",
        status: "ตรวจสอบเอกสาร",
        appliedDate: "15 ม.ค. 2568",
        email: "somchai.j@email.com",
        phone: "0812345678"
    },
    {
        id: 2,
        year: 2568,
        semester: 1,
        studentName: "สมหญิง รักเรียน",
        jobTitle: "Backend Developer (In-tern)",
        status: "นัดสัมภาษณ์",
        appliedDate: "14 ม.ค. 2568",
        email: "somying.r@email.com",
        phone: "0823456789"
    },
    {
        id: 3,
        year: 2568,
        semester: 1,
        studentName: "วิชัย มีความสุข",
        jobTitle: "Mobile Developer",
        status: "รอพิจารณา",
        appliedDate: "13 ม.ค. 2568",
        email: "wichai.m@email.com",
        phone: "0834567890"
    },
    {
        id: 4,
        year: 2568,
        semester: 1,
        studentName: "สุดา ขยัน",
        jobTitle: "Data Analytics",
        status: "ผ่านการคัดเลือก",
        appliedDate: "12 ม.ค. 2568",
        email: "suda.k@email.com",
        phone: "0845678901"
    },
    {
        id: 5,
        year: 2568,
        semester: 1,
        studentName: "ประเสริฐ ดีเด่น",
        jobTitle: "UI/UX Designer",
        status: "ไม่ผ่านการคัดเลือก",
        appliedDate: "11 ม.ค. 2568",
        email: "prasert.d@email.com",
        phone: "0856789012"
    },
    {
        id: 6,
        year: 2568,
        semester: 2,
        studentName: "มานี สดใส",
        jobTitle: "Frontend Developer (In-tern)",
        status: "ตรวจสอบเอกสาร",
        appliedDate: "10 ม.ค. 2568",
        email: "manee.s@email.com",
        phone: "0867890123"
    },
    {
        id: 7,
        year: 2568,
        semester: 2,
        studentName: "จิรายุ เก่งกล้า",
        jobTitle: "Backend Developer (In-tern)",
        status: "นัดสัมภาษณ์",
        appliedDate: "09 ม.ค. 2568",
        email: "jirayu.k@email.com",
        phone: "0878901234"
    }
]);

// =======================================================
// COMPUTED PROPERTIES
// =======================================================
const tabs = computed(() => [
    { value: 'all', label: 'ทั้งหมด', count: mockApplicationsData.value.length },
    { value: 'ตรวจสอบเอกสาร', label: 'ตรวจสอบเอกสาร', count: reviewingCount.value },
    { value: 'นัดสัมภาษณ์', label: 'นัดสัมภาษณ์', count: interviewCount.value },
    { value: 'รอพิจารณา', label: 'รอพิจารณา', count: pendingCount.value },
    { value: 'ผ่านการคัดเลือก', label: 'ผ่านการคัดเลือก', count: approvedCount.value },
    { value: 'ไม่ผ่านการคัดเลือก', label: 'ไม่ผ่านการคัดเลือก', count: rejectedCount.value }
]);

const reviewingCount = computed(() =>
    mockApplicationsData.value.filter(a => a.status === 'ตรวจสอบเอกสาร').length
);

const interviewCount = computed(() =>
    mockApplicationsData.value.filter(a => a.status === 'นัดสัมภาษณ์').length
);

const pendingCount = computed(() =>
    mockApplicationsData.value.filter(a => a.status === 'รอพิจารณา').length
);

const approvedCount = computed(() =>
    mockApplicationsData.value.filter(a => a.status === 'ผ่านการคัดเลือก').length
);

const rejectedCount = computed(() =>
    mockApplicationsData.value.filter(a => a.status === 'ไม่ผ่านการคัดเลือก').length
);

const totalCount = computed(() => mockApplicationsData.value.length);

const filteredApplications = computed(() => {
    let result = [...mockApplicationsData.value];

    // Filter by tab
    if (activeTab.value !== 'all') {
        result = result.filter(app => app.status === activeTab.value);
    }

    // Filter by year
    result = result.filter(app => app.year === selectedYear.value);

    // Filter by semester
    result = result.filter(app => app.semester === selectedSemester.value);

    // Filter by job
    if (selectedJob.value !== 'ทั้งหมด') {
        result = result.filter(app => app.jobTitle === selectedJob.value);
    }

    // Filter by search
    if (searchText.value) {
        const search = searchText.value.toLowerCase();
        result = result.filter(app =>
            app.studentName.toLowerCase().includes(search) ||
            app.jobTitle.toLowerCase().includes(search)
        );
    }

    // Sort
    result.sort((a, b) => {
        const dateA = new Date(a.appliedDate.split(' ').reverse().join('-'));
        const dateB = new Date(b.appliedDate.split(' ').reverse().join('-'));
        return sortAscending.value ? dateA - dateB : dateB - dateA;
    });

    return result;
});

const totalPages = computed(() =>
    Math.ceil(filteredApplications.value.length / itemsPerPage.value)
);

const paginatedApplications = computed(() => {
    const start = (currentPage.value - 1) * itemsPerPage.value;
    const end = start + itemsPerPage.value;
    return filteredApplications.value.slice(start, end);
});

const visiblePages = computed(() => {
    const pages = [];
    const total = totalPages.value;
    const current = currentPage.value;

    if (total <= 7) {
        for (let i = 1; i <= total; i++) {
            pages.push(i);
        }
    } else {
        if (current <= 4) {
            for (let i = 1; i <= 5; i++) pages.push(i);
            pages.push('...');
            pages.push(total);
        } else if (current >= total - 3) {
            pages.push(1);
            pages.push('...');
            for (let i = total - 4; i <= total; i++) pages.push(i);
        } else {
            pages.push(1);
            pages.push('...');
            for (let i = current - 1; i <= current + 1; i++) pages.push(i);
            pages.push('...');
            pages.push(total);
        }
    }

    return pages;
});

// =======================================================
// METHODS
// =======================================================
const getStatusClass = (status) => {
    const statusMap = {
        'ตรวจสอบเอกสาร': 'reviewing',
        'นัดสัมภาษณ์': 'interview',
        'รอพิจารณา': 'pending',
        'ผ่านการคัดเลือก': 'approved',
        'ไม่ผ่านการคัดเลือก': 'rejected'
    };
    return statusMap[status] || '';
};

const getStatusBackgroundClass = (status) => {
    const bgMap = {
        'ตรวจสอบเอกสาร': 'bg-reviewing',
        'นัดสัมภาษณ์': 'bg-interview',
        'รอพิจารณา': 'bg-pending',
        'ผ่านการคัดเลือก': 'bg-approved',
        'ไม่ผ่านการคัดเลือก': 'bg-rejected'
    };
    return bgMap[status] || '';
};

const getEmptyMessage = () => {
    if (searchText.value) {
        return `ไม่พบใบสมัครงานที่ตรงกับคำค้นหา "${searchText.value}"`;
    }
    if (activeTab.value === 'ตรวจสอบเอกสาร') {
        return 'ยังไม่มีใบสมัครงานที่รอตรวจสอบเอกสาร';
    }
    if (activeTab.value === 'นัดสัมภาษณ์') {
        return 'ยังไม่มีใบสมัครงานที่นัดสัมภาษณ์';
    }
    if (activeTab.value === 'รอพิจารณา') {
        return 'ยังไม่มีใบสมัครงานที่รอพิจารณา';
    }
    if (activeTab.value === 'ผ่านการคัดเลือก') {
        return 'ยังไม่มีใบสมัครงานที่ผ่านการคัดเลือก';
    }
    if (activeTab.value === 'ไม่ผ่านการคัดเลือก') {
        return 'ยังไม่มีใบสมัครงานที่ไม่ผ่านการคัดเลือก';
    }
    return 'ยังไม่มีใบสมัครงานในระบบ';
};

const toggleSort = () => {
    sortAscending.value = !sortAscending.value;
};

const selectYear = (year) => {
    selectedYear.value = year;
    isYearDropdownOpen.value = false;
    currentPage.value = 1;
};

const selectJob = (job) => {
    selectedJob.value = job;
    isJobDropdownOpen.value = false;
    currentPage.value = 1;
};

const handleSearch = () => {
    currentPage.value = 1;
};

const toggleRow = (id) => {
    if (expandedRows.value.has(id)) {
        expandedRows.value.delete(id);
    } else {
        expandedRows.value.add(id);
    }
};

const updateStatus = (id, newStatus) => {
    const app = mockApplicationsData.value.find(a => a.id === id);
    if (app) {
        app.status = newStatus;
        console.log(`Updated application ${id} to status: ${newStatus}`);
    }
};

const viewApplication = (id) => {
    console.log('View application:', id);
    // TODO: Navigate to application details page
};

const handleClickOutside = (event) => {
    if (yearDropdownRef.value && !yearDropdownRef.value.contains(event.target)) {
        isYearDropdownOpen.value = false;
    }
    if (jobDropdownRef.value && !jobDropdownRef.value.contains(event.target)) {
        isJobDropdownOpen.value = false;
    }
};

// =======================================================
// WATCHERS
// =======================================================
watch([activeTab, searchText, selectedYear, selectedSemester, selectedJob], () => {
    currentPage.value = 1;
    expandedRows.value.clear();
});

// =======================================================
// LIFECYCLE
// =======================================================
onMounted(() => {
    document.addEventListener('click', handleClickOutside);
    isLoading.value = true;
    setTimeout(() => {
        isLoading.value = false;
    }, 500);
});

onBeforeUnmount(() => {
    document.removeEventListener('click', handleClickOutside);
});
</script>

<style scoped>
@import url("https://fonts.googleapis.com/css2?family=Kanit:wght@400;500;600;700&display=swap");

* {
    font-family: "Kanit", sans-serif;
    box-sizing: border-box;
}

.job-applications-page {
    padding: 20px;
    max-width: 1400px;
    margin: 0 auto;
}

/* =================================== */
/* PAGE HEADER */
/* =================================== */
.page-header {
    display: flex;
    justify-content: space-between;
    align-items: flex-start;
    margin-bottom: 30px;
    gap: 20px;
}

.header-content h1 {
    font-size: 28px;
    font-weight: 700;
    color: #222;
    margin: 0 0 8px 0;
}

.page-subtitle {
    font-size: 15px;
    color: #666;
    margin: 0;
    line-height: 1.6;
}

/* =================================== */
/* STATS OVERVIEW */
/* =================================== */
.stats-overview {
    display: grid;
    grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
    gap: 20px;
    margin-bottom: 30px;
}

.stat-card {
    background: #fff;
    border-radius: 16px;
    padding: 24px;
    display: flex;
    align-items: center;
    gap: 16px;
    box-shadow: 0 2px 12px rgba(3, 114, 102, 0.07);
    border: 1px solid #eef0ef;
    transition: all 0.2s;
}

.stat-card:hover {
    transform: translateY(-2px);
    box-shadow: 0 4px 20px rgba(3, 114, 102, 0.12);
}

.stat-icon {
    width: 56px;
    height: 56px;
    border-radius: 14px;
    display: flex;
    align-items: center;
    justify-content: center;
    flex-shrink: 0;
}

.stat-icon.reviewing {
    background: linear-gradient(135deg, #3b82f6 0%, #2563eb 100%);
    color: #fff;
}

.stat-icon.interview {
    background: linear-gradient(135deg, #06b6d4 0%, #0891b2 100%);
    color: #fff;
}

.stat-icon.pending {
    background: linear-gradient(135deg, #f59e0b 0%, #d97706 100%);
    color: #fff;
}

.stat-icon.approved {
    background: linear-gradient(135deg, #10b981 0%, #059669 100%);
    color: #fff;
}

.stat-icon.rejected {
    background: linear-gradient(135deg, #ef4444 0%, #dc2626 100%);
    color: #fff;
}

.stat-icon.total {
    background: linear-gradient(135deg, #037266 0%, #026b5b 100%);
    color: #fff;
}

.stat-info {
    display: flex;
    flex-direction: column;
    gap: 4px;
}

.stat-label {
    font-size: 14px;
    color: #666;
    font-weight: 500;
}

.stat-value {
    font-size: 32px;
    font-weight: 700;
    color: #222;
    line-height: 1;
}

/* =================================== */
/* FILTER SECTION */
/* =================================== */
.filter-section {
    background: #fff;
    border-radius: 16px;
    padding: 20px;
    margin-bottom: 25px;
    box-shadow: 0 2px 12px rgba(3, 114, 102, 0.07);
    border: 1px solid #eef0ef;
}

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
    font-weight: 500;
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

.filters-row {
    display: flex;
    gap: 20px;
    margin-bottom: 20px;
    flex-wrap: wrap;
    align-items: flex-end;
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
    flex-wrap: wrap;
}

.search-box {
    flex: 1;
    min-width: 250px;
    position: relative;
    display: flex;
    align-items: center;
}

.search-icon {
    position: absolute;
    left: 14px;
    color: #999;
}

.search-box input {
    width: 100%;
    padding: 11px 14px 11px 42px;
    border: 2px solid #e5e5e5;
    border-radius: 10px;
    font-size: 14px;
    transition: all 0.2s;
    font-family: "Kanit", sans-serif;
}

.search-box input:focus {
    outline: none;
    border-color: #037266;
    background: #f0f8f7;
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
    font-size: 14px;
    transition: all 0.2s;
    height: 45px;
    width: 90px;
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
    height: 45px;
    width: 70px;
}

.btn-sort:hover {
    background: #e0e0e0;
}

/* =================================== */
/* LOADING & EMPTY STATE */
/* =================================== */
.loading-state {
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    padding: 80px 20px;
    gap: 20px;
}

.spinner {
    width: 48px;
    height: 48px;
    border: 4px solid #f0f0f0;
    border-top-color: #037266;
    border-radius: 50%;
    animation: spin 0.8s linear infinite;
}

@keyframes spin {
    to {
        transform: rotate(360deg);
    }
}

.loading-state p {
    font-size: 15px;
    color: #666;
    margin: 0;
}

.empty-state {
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    padding: 80px 20px;
    gap: 16px;
}

.empty-state svg {
    color: #d0d0d0;
}

.empty-state h3 {
    font-size: 20px;
    font-weight: 600;
    color: #555;
    margin: 0;
}

.empty-state p {
    font-size: 15px;
    color: #888;
    margin: 0;
    text-align: center;
    max-width: 400px;
}

/* =================================== */
/* TABLE */
/* =================================== */
.table-wrapper {
    background: #fff;
    border-radius: 16px;
    overflow: hidden;
    box-shadow: 0 2px 12px rgba(3, 114, 102, 0.07);
    border: 1px solid #eef0ef;
}

.table-container {
    overflow-x: auto;
}

.applications-table {
    width: 100%;
    border-collapse: collapse;
    min-width: 800px;
}

.applications-table thead {
    background: #f7f9f8;
}

.applications-table th {
    padding: 16px 20px;
    text-align: left;
    font-size: 14px;
    font-weight: 700;
    color: #555;
    text-transform: uppercase;
    letter-spacing: 0.5px;
    border-bottom: 2px solid #e5e5e5;
}

.applications-table tbody tr.clickable-row {
    border-bottom: 1px solid #f0f0f0;
    transition: all 0.2s ease;
    cursor: pointer;
}

.applications-table tbody tr.clickable-row:hover {
    background: #fafbfb;
    box-shadow: 0 2px 8px rgba(3, 114, 102, 0.04);
}

.applications-table tbody tr.clickable-row.expanded {
    background: #f8fafa;
    border-bottom: none;
}

.applications-table td {
    padding: 18px 20px;
    font-size: 14px;
    color: #2d3748;
    vertical-align: middle;
}

.expand-cell {
    width: 50px;
    padding: 18px 10px 18px 20px !important;
}

.expand-icon {
    transition: transform 0.3s ease;
    color: #9ca3af;
}

.expand-icon.rotated {
    transform: rotate(90deg);
    color: #037266;
}

.applicant-name {
    font-weight: 600;
    color: #1a202c;
}

.job-position {
    color: #4a5568;
    font-size: 13px;
}

.status-badge {
    display: inline-flex;
    align-items: center;
    gap: 6px;
    padding: 6px 12px;
    border-radius: 20px;
    font-size: 13px;
    font-weight: 600;
    white-space: nowrap;
}

.status-badge.reviewing {
    background: #dbeafe;
    color: #1e40af;
}

.status-badge.interview {
    background: #e0f2fe;
    color: #0369a1;
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

.status-dot {
    width: 8px;
    height: 8px;
    border-radius: 50%;
    background: currentColor;
}

.date {
    color: #4a5568;
    font-weight: 500;
}

/* =================================== */
/* DETAILS ROW */
/* =================================== */
.details-row {
    border-bottom: 1px solid #e5e7e6 !important;
}

.details-row.bg-reviewing {
    background: #f0f9ff;
}

.details-row.bg-interview {
    background: #ecfeff;
}

.details-row.bg-pending {
    background: #fffbeb;
}

.details-row.bg-approved {
    background: #f0fdf4;
}

.details-row.bg-rejected {
    background: #fef2f2;
}

.details-content {
    padding: 24px 30px 24px 70px;
    display: flex;
    justify-content: space-between;
    align-items: flex-start;
    gap: 30px;
    animation: slideDown 0.3s ease-out;
}

@keyframes slideDown {
    from {
        opacity: 0;
        transform: translateY(-10px);
    }

    to {
        opacity: 1;
        transform: translateY(0);
    }
}

.student-info-section {
    flex: 1;
}

.info-header {
    margin-bottom: 16px;
}

.info-header h4 {
    font-size: 14px;
    font-weight: 600;
    color: #6b7280;
    margin: 0;
    letter-spacing: 0;
}

.student-name {
    font-size: 20px;
    font-weight: 700;
    color: #1a202c;
    margin-bottom: 16px;
}

.contact-list {
    display: flex;
    flex-direction: column;
    gap: 8px;
}

.contact-item {
    display: flex;
    align-items: center;
    gap: 8px;
    font-size: 14px;
    color: #4a5568;
}

.contact-icon {
    color: #9ca3af;
    flex-shrink: 0;
}

/* =================================== */
/* ACTIONS SECTION */
/* =================================== */
.actions-section {
    display: flex;
    flex-direction: column;
    gap: 10px;
    min-width: 200px;
}

.btn-status {
    padding: 11px 20px;
    border: none;
    border-radius: 10px;
    cursor: pointer;
    font-size: 14px;
    font-weight: 600;
    transition: all 0.2s ease;
    white-space: nowrap;
    width: 100%;
    text-align: center;
    display: flex;
    align-items: center;
    justify-content: center;
    gap: 6px;
}

.btn-status.btn-approve {
    background: #d1fae5;
    color: #065f46;
    border: 2px solid transparent;
}

.btn-status.btn-approve:hover {
    background: #a7f3d0;
    transform: translateY(-2px);
    box-shadow: 0 4px 12px rgba(6, 95, 70, 0.2);
}

.btn-status.btn-interview-done {
    background: #e0f2fe;
    color: #0369a1;
    border: 2px solid transparent;
}

.btn-status.btn-interview-done:hover {
    background: #bae6fd;
    transform: translateY(-2px);
    box-shadow: 0 4px 12px rgba(3, 105, 161, 0.2);
}

.btn-status.btn-reject {
    background: #fee2e2;
    color: #991b1b;
    border: 2px solid transparent;
}

.btn-status.btn-reject:hover {
    background: #fecaca;
    transform: translateY(-2px);
    box-shadow: 0 4px 12px rgba(153, 27, 27, 0.2);
}

.btn-status.btn-view-application {
    background: #dbeafe;
    color: #1e40af;
    border: 2px solid transparent;
}

.btn-status.btn-view-application:hover {
    background: #bfdbfe;
    transform: translateY(-2px);
    box-shadow: 0 4px 12px rgba(30, 64, 175, 0.2);
}

/* =================================== */
/* PAGINATION */
/* =================================== */
.pagination {
    display: flex;
    justify-content: center;
    align-items: center;
    gap: 8px;
    margin-top: 30px;
    padding: 20px;
}

.btn-page {
    min-width: 40px;
    height: 40px;
    padding: 0 12px;
    border: 2px solid #e5e5e5;
    background: #fff;
    border-radius: 10px;
    font-size: 14px;
    font-weight: 600;
    color: #555;
    cursor: pointer;
    transition: all 0.2s;
    display: flex;
    align-items: center;
    justify-content: center;
}

.btn-page:hover:not(:disabled) {
    border-color: #037266;
    color: #037266;
    background: #f0f8f7;
}

.btn-page.active {
    background: linear-gradient(135deg, #037266 0%, #026b5b 100%);
    color: #fff;
    border-color: #037266;
}

.btn-page:disabled {
    opacity: 0.4;
    cursor: not-allowed;
}

.page-ellipsis {
    padding: 0 8px;
    color: #999;
    font-weight: 600;
}

/* =================================== */
/* RESPONSIVE */
/* =================================== */
@media (max-width: 1200px) {
    .details-content {
        flex-direction: column;
        gap: 24px;
    }

    .actions-section {
        width: 100%;
        flex-direction: row;
        flex-wrap: wrap;
    }

    .btn-status {
        flex: 1;
        min-width: 160px;
    }
}

@media (max-width: 1024px) {
    .stats-overview {
        grid-template-columns: repeat(3, 1fr);
    }
}

@media (max-width: 900px) {
    .filters-row {
        gap: 16px;
    }

    .custom-dropdown {
        width: 220px;
    }

    .applications-table {
        font-size: 13px;
    }

    .applications-table th,
    .applications-table td {
        padding: 14px 16px;
    }

    .details-content {
        padding: 20px 24px 20px 50px;
    }
}

@media (max-width: 768px) {
    .job-applications-page {
        padding: 15px;
    }

    .page-header {
        flex-direction: column;
        align-items: stretch;
    }

    .stats-overview {
        grid-template-columns: repeat(2, 1fr);
    }

    .search-sort {
        flex-direction: column;
        align-items: stretch;
    }

    .search-box {
        min-width: 100%;
    }

    .btn-search,
    .btn-sort {
        width: 100%;
        justify-content: center;
    }

    .applications-table-container {
        overflow-x: auto;
    }

    .applications-table {
        min-width: 800px;
    }

    .details-content {
        margin-left: 10px;
        padding: 16px 20px;
    }
}

@media (max-width: 650px) {
    .filters-row {
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

    .applications-table th,
    .applications-table td {
        padding: 12px;
        font-size: 12px;
    }

    .status-badge {
        font-size: 11px;
        padding: 4px 10px;
    }

    .expand-cell {
        padding: 12px 8px !important;
    }

    .details-content {
        padding: 14px 16px;
        margin-left: 5px;
    }

    .info-header h4 {
        font-size: 14px;
    }

    .btn-status {
        font-size: 12px;
        padding: 8px 14px;
        min-width: 120px;
    }
}

@media (max-width: 480px) {
    .stats-overview {
        grid-template-columns: 1fr;
    }

    .header-content h1 {
        font-size: 24px;
    }

    .tabs {
        gap: 4px;
        overflow-x: auto;
        -webkit-overflow-scrolling: touch;
        scrollbar-width: none;
    }

    .tabs::-webkit-scrollbar {
        display: none;
    }

    .tab-btn {
        padding: 8px 12px;
        font-size: 13px;
        white-space: nowrap;
    }

    .pagination {
        gap: 4px;
    }

    .btn-page {
        min-width: 36px;
        height: 36px;
        font-size: 13px;
    }

    .actions-section {
        flex-direction: column;
    }

    .btn-status {
        width: 100%;
        min-width: unset;
    }
}
</style>