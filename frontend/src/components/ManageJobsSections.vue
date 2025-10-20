<!-- frontend/src/components/ManageJobsSections.vue -->
<template>
    <div class="manage-jobs-page">
        <div class="page-header">
            <div class="header-content">
                <h1>จัดการประกาศงาน</h1>
                <p class="page-subtitle">จัดการ แก้ไข และติดตามสถานะของประกาศงานทั้งหมด</p>
            </div>
            <button class="btn-create" @click="createNewJob">
                <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none"
                    stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                    <line x1="12" y1="5" x2="12" y2="19"></line>
                    <line x1="5" y1="12" x2="19" y2="12"></line>
                </svg>
                สร้างประกาศงานใหม่
            </button>
        </div>

        <div class="stats-overview">
            <div class="stat-card">
                <div class="stat-icon approved">
                    <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none"
                        stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                        <path d="M22 11.08V12a10 10 0 1 1-5.93-9.14"></path>
                        <polyline points="22 4 12 14.01 9 11.01"></polyline>
                    </svg>
                </div>
                <div class="stat-info">
                    <span class="stat-label">ผ่านการตรวจสอบ</span>
                    <span class="stat-value">{{ approvedCount }}</span>
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
                    <span class="stat-label">รอตรวจสอบ</span>
                    <span class="stat-value">{{ pendingCount }}</span>
                </div>
            </div>

            <div class="stat-card">
                <div class="stat-icon rejected">
                    <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none"
                        stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                        <circle cx="12" cy="12" r="10"></circle>
                        <line x1="15" y1="9" x2="9" y2="15"></line>
                        <line x1="9" y1="9" x2="15" y2="15"></line>
                    </svg>
                </div>
                <div class="stat-info">
                    <span class="stat-label">ไม่ผ่านการตรวจสอบ</span>
                    <span class="stat-value">{{ rejectedCount }}</span>
                </div>
            </div>

            <div class="stat-card">
                <div class="stat-icon total">
                    <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none"
                        stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                        <path d="M13 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V9z"></path>
                        <polyline points="13 2 13 9 20 9"></polyline>
                    </svg>
                </div>
                <div class="stat-info">
                    <span class="stat-label">ทั้งหมด</span>
                    <span class="stat-value">{{ totalCount }}</span>
                </div>
            </div>
        </div>

        <div class="filter-section">
            <div class="tabs">
                <button v-for="tab in tabs" :key="tab.value" class="tab-btn"
                    :class="{ active: activeTab === tab.value }" @click="activeTab = tab.value">
                    {{ tab.label }}
                </button>
            </div>

            <div class="search-sort">
                <div class="search-box">
                    <svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none"
                        stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                        <circle cx="11" cy="11" r="8"></circle>
                        <path d="m21 21-4.35-4.35"></path>
                    </svg>
                    <input type="text" v-model="searchText" placeholder="ค้นหาชื่อตำแหน่งงาน..." />
                </div>
                <button class="btn-search" @click="handleSearch">ค้นหา</button>

                <div class="filter-dropdown" ref="yearFilterRef">
                    <button class="btn-filter-extra" @click="isYearFilterOpen = !isYearFilterOpen">
                        <svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24"
                            fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round"
                            stroke-linejoin="round">
                            <rect x="3" y="4" width="18" height="18" rx="2" ry="2" />
                            <line x1="16" y1="2" x2="16" y2="6" />
                            <line x1="8" y1="2" x2="8" y2="6" />
                            <line x1="3" y1="10" x2="21" y2="10" />
                        </svg>
                        ปีการศึกษา
                        <span v-if="selectedYear !== 'ทั้งหมด'" class="filter-count">1</span>
                    </button>
                    <transition name="dropdown-fade">
                        <div v-if="isYearFilterOpen" class="filter-menu">
                            <label v-for="year in academicYears" :key="year" class="filter-radio">
                                <input type="radio" :value="year" v-model="selectedYear" />
                                <span>{{ year }}</span>
                            </label>
                        </div>
                    </transition>
                </div>

                <div class="filter-dropdown" ref="semesterFilterRef">
                    <button class="btn-filter-extra" @click="isSemesterFilterOpen = !isSemesterFilterOpen">
                        <svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24"
                            fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round"
                            stroke-linejoin="round">
                            <path d="M4 19.5A2.5 2.5 0 0 1 6.5 17H20"></path>
                            <path d="M6.5 2H20v20H6.5A2.5 2.5 0 0 1 4 19.5v-15A2.5 2.5 0 0 1 6.5 2z"></path>
                        </svg>
                        ภาคเรียน
                        <span v-if="selectedSemester !== 'ทั้งหมด'" class="filter-count">1</span>
                    </button>
                    <transition name="dropdown-fade">
                        <div v-if="isSemesterFilterOpen" class="filter-menu">
                            <label v-for="sem in semesters" :key="sem" class="filter-radio">
                                <input type="radio" :value="sem" v-model="selectedSemester" />
                                <span>{{ sem }}</span>
                            </label>
                        </div>
                    </transition>
                </div>

                <div class="filter-dropdown" ref="departmentFilterRef">
                    <button class="btn-filter-extra" @click="isDepartmentFilterOpen = !isDepartmentFilterOpen">
                        <svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24"
                            fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round"
                            stroke-linejoin="round">
                            <polygon points="22 3 2 3 10 12.46 10 19 14 21 14 12.46 22 3"></polygon>
                        </svg>
                        ภาควิชา
                        <span v-if="selectedDepartment !== 'ทั้งหมด'" class="filter-count">1</span>
                    </button>
                    <transition name="dropdown-fade">
                        <div v-if="isDepartmentFilterOpen" class="filter-menu">
                            <label v-for="dept in departments" :key="dept" class="filter-radio">
                                <input type="radio" :value="dept" v-model="selectedDepartment" />
                                <span>{{ dept }}</span>
                            </label>
                        </div>
                    </transition>
                </div>

                <button @click="toggleSort" class="btn-sort"
                    :title="sortAscending ? 'เรียงจากเก่าสุด' : 'เรียงจากใหม่สุด'">
                    <svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" fill="none" stroke="#9CA3AF"
                        stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                        <path d="M3 6h18M3 12h12M3 18h6" />
                    </svg>
                </button>

                <button v-if="hasActiveFilters" class="btn-clear-all" @click="clearAllFilters">
                    <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none"
                        stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                        <line x1="18" y1="6" x2="6" y2="18"></line>
                        <line x1="6" y1="6" x2="18" y2="18"></line>
                    </svg>
                    ล้าง
                </button>
            </div>
        </div>

        <div v-if="isLoading" class="loading-state">
            <div class="spinner"></div>
            <p>กำลังโหลดข้อมูล...</p>
        </div>

        <div v-else-if="filteredJobs.length === 0" class="empty-state">
            <svg xmlns="http://www.w3.org/2000/svg" width="64" height="64" viewBox="0 0 24 24" fill="none"
                stroke="currentColor" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round">
                <rect x="3" y="3" width="18" height="18" rx="2" ry="2"></rect>
                <line x1="9" y1="9" x2="15" y2="9"></line>
                <line x1="9" y1="15" x2="15" y2="15"></line>
            </svg>
            <h3>ไม่พบประกาศงาน</h3>
            <p>{{ getEmptyMessage() }}</p>
            <button class="btn-create-empty" @click="createNewJob">
                <svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none"
                    stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                    <line x1="12" y1="5" x2="12" y2="19"></line>
                    <line x1="5" y1="12" x2="19" y2="12"></line>
                </svg>
                สร้างประกาศงานใหม่
            </button>
        </div>

        <div v-else class="jobs-table-container">
            <table class="jobs-table">
                <thead>
                    <tr>
                        <th class="col-checkbox">
                            <input type="checkbox" v-model="selectAll" @change="toggleSelectAll" />
                        </th>
                        <th class="col-title">ชื่อตำแหน่งงาน</th>
                        <th class="col-department">ภาควิชา</th>
                        <th class="col-status">สถานะ</th>
                        <th class="col-duration">ระยะเวลารับสมัคร</th>
                        <th class="col-date">วันที่ประกาศ</th>
                        <th class="col-applicants">ผู้สมัคร</th>
                        <th class="col-actions">จัดการ</th>
                    </tr>
                </thead>
                <tbody>
                    <tr v-for="job in paginatedJobs" :key="job.id" class="job-row"
                        :class="{ selected: selectedJobs.includes(job.id) }">
                        <td class="col-checkbox">
                            <input type="checkbox" :value="job.id" v-model="selectedJobs" />
                        </td>
                        <td class="col-title">
                            <div class="title-cell">
                                <div class="title-info">
                                    <h4>{{ job.title }}</h4>
                                    <p>{{ job.year }} / {{ job.semester }}</p>
                                </div>
                            </div>
                        </td>
                        <td class="col-department">
                            <span class="department-text">{{ job.department }}</span>
                        </td>
                        <td class="col-status">
                            <span class="status-badge" :class="getStatusClass(job.status)">
                                <span class="status-dot"></span>
                                {{ job.status }}
                            </span>
                        </td>
                        <td class="col-duration">
                            <div class="duration-cell">
                                <svg xmlns="http://www.w3.org/2000/svg" width="14" height="14" viewBox="0 0 24 24"
                                    fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round"
                                    stroke-linejoin="round">
                                    <circle cx="12" cy="12" r="10"></circle>
                                    <polyline points="12 6 12 12 16 14"></polyline>
                                </svg>
                                {{ job.duration }} วัน
                            </div>
                        </td>
                        <td class="col-date">
                            <div class="date-cell">
                                <svg xmlns="http://www.w3.org/2000/svg" width="14" height="14" viewBox="0 0 24 24"
                                    fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round"
                                    stroke-linejoin="round">
                                    <rect x="3" y="4" width="18" height="18" rx="2" ry="2" />
                                    <line x1="16" y1="2" x2="16" y2="6" />
                                    <line x1="8" y1="2" x2="8" y2="6" />
                                    <line x1="3" y1="10" x2="21" y2="10" />
                                </svg>
                                {{ job.postedDate || '-' }}
                            </div>
                        </td>
                        <td class="col-applicants">
                            <div class="applicants-cell">
                                <svg xmlns="http://www.w3.org/2000/svg" width="14" height="14" viewBox="0 0 24 24"
                                    fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round"
                                    stroke-linejoin="round">
                                    <path d="M17 21v-2a4 4 0 0 0-4-4H5a4 4 0 0 0-4 4v2"></path>
                                    <circle cx="9" cy="7" r="4"></circle>
                                    <path d="M23 21v-2a4 4 0 0 0-3-3.87"></path>
                                    <path d="M16 3.13a4 4 0 0 1 0 7.75"></path>
                                </svg>
                                {{ job.applicants }} คน
                            </div>
                        </td>
                        <td class="col-actions">
                            <div class="action-buttons">
                                <button class="btn-action btn-view" @click="openSidePanel(job)"
                                    title="ดูรายละเอียด">
                                    <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16"
                                        viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"
                                        stroke-linecap="round" stroke-linejoin="round">
                                        <path d="M1 12s4-8 11-8 11 8 11 8-4 8-11 8-11-8-11-8z"></path>
                                        <circle cx="12" cy="12" r="3"></circle>
                                    </svg>
                                </button>
                                <button class="btn-action btn-edit" @click="editJob(job)" title="แก้ไข">
                                    <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16"
                                        viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"
                                        stroke-linecap="round" stroke-linejoin="round">
                                        <path d="M11 4H4a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2v-7">
                                        </path>
                                        <path d="M18.5 2.5a2.121 2.121 0 0 1 3 3L12 15l-4 1 1-4 9.5-9.5z">
                                        </path>
                                    </svg>
                                </button>
                                <button class="btn-action btn-delete" @click="confirmDelete(job)" title="ลบ">
                                    <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16"
                                        viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"
                                        stroke-linecap="round" stroke-linejoin="round">
                                        <polyline points="3 6 5 6 21 6"></polyline>
                                        <path
                                            d="M19 6v14a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2V6m3 0V4a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v2">
                                        </path>
                                        <line x1="10" y1="11" x2="10" y2="17"></line>
                                        <line x1="14" y1="11" x2="14" y2="17"></line>
                                    </svg>
                                </button>
                            </div>
                        </td>
                    </tr>
                </tbody>
            </table>
        </div>

        <div v-if="filteredJobs.length > 0" class="bulk-actions-bar" :class="{ visible: selectedJobs.length > 0 }">
            <div class="bulk-info">
                <span>เลือกแล้ว {{ selectedJobs.length }} รายการ</span>
            </div>
            <div class="bulk-buttons">
                <button class="btn-bulk btn-delete-bulk" @click="bulkDelete">
                    <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none"
                        stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                        <polyline points="3 6 5 6 21 6"></polyline>
                        <path d="M19 6v14a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2V6m3 0V4a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v2">
                        </path>
                    </svg>
                    ลบ
                </button>
            </div>
        </div>

        <div v-if="filteredJobs.length > itemsPerPage" class="pagination">
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

        <!-- View Job Detail Side Panel -->
        <div v-if="selectedJobDetail" class="side-panel-overlay" @click.self="closeSidePanel">
            <div class="side-panel">
                <button class="side-close" @click="closeSidePanel">&times;</button>

                <div class="side-panel-inner">
                    <div class="side-header">
                        <h2>{{ selectedJobDetail.title }}</h2>
                        <div class="job-meta-inline">
                            <span class="department-text-meta">{{ selectedJobDetail.department }}</span>
                            <span class="status-badge" :class="getStatusClass(selectedJobDetail.status)">
                                <span class="status-dot"></span>
                                {{ selectedJobDetail.status }}
                            </span>
                        </div>
                        <div class="job-details-grid">
                            <div class="detail-item">
                                <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" width="14" height="14"
                                    fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round"
                                    stroke-linejoin="round">
                                    <rect x="3" y="4" width="18" height="18" rx="2" ry="2" />
                                    <line x1="16" y1="2" x2="16" y2="6" />
                                    <line x1="8" y1="2" x2="8" y2="6" />
                                    <line x1="3" y1="10" x2="21" y2="10" />
                                </svg>
                                <span>{{ selectedJobDetail.year }} / {{ selectedJobDetail.semester }}</span>
                            </div>
                            <div class="detail-item">
                                <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" width="14" height="14"
                                    fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round"
                                    stroke-linejoin="round">
                                    <circle cx="12" cy="12" r="10"></circle>
                                    <polyline points="12 6 12 12 16 14"></polyline>
                                </svg>
                                <span>{{ selectedJobDetail.duration }} วัน</span>
                            </div>
                            <div class="detail-item">
                                <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" width="14" height="14"
                                    fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round"
                                    stroke-linejoin="round">
                                    <path d="M17 21v-2a4 4 0 0 0-4-4H5a4 4 0 0 0-4 4v2"></path>
                                    <circle cx="9" cy="7" r="4"></circle>
                                    <path d="M23 21v-2a4 4 0 0 0-3-3.87"></path>
                                    <path d="M16 3.13a4 4 0 0 1 0 7.75"></path>
                                </svg>
                                <span>{{ selectedJobDetail.applicants }} คน</span>
                            </div>
                            <div class="detail-item" v-if="selectedJobDetail.postedDate">
                                <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" width="14" height="14"
                                    fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round"
                                    stroke-linejoin="round">
                                    <rect x="3" y="4" width="18" height="18" rx="2" ry="2" />
                                    <line x1="16" y1="2" x2="16" y2="6" />
                                    <line x1="8" y1="2" x2="8" y2="6" />
                                    <line x1="3" y1="10" x2="21" y2="10" />
                                </svg>
                                <span>ประกาศ: {{ selectedJobDetail.postedDate }}</span>
                            </div>
                        </div>
                    </div>

                    <div class="side-action-buttons">
                        <button class="btn-side-action btn-edit" @click="editJob(selectedJobDetail)">
                            <svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24"
                                fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round"
                                stroke-linejoin="round">
                                <path d="M11 4H4a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2v-7"></path>
                                <path d="M18.5 2.5a2.121 2.121 0 0 1 3 3L12 15l-4 1 1-4 9.5-9.5z"></path>
                            </svg>
                            <span>แก้ไข</span>
                        </button>
                        <button class="btn-side-action btn-delete" @click="confirmDelete(selectedJobDetail)">
                            <svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24"
                                fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round"
                                stroke-linejoin="round">
                                <polyline points="3 6 5 6 21 6"></polyline>
                                <path
                                    d="M19 6v14a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2V6m3 0V4a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v2">
                                </path>
                                <line x1="10" y1="11" x2="10" y2="17"></line>
                                <line x1="14" y1="11" x2="14" y2="17"></line>
                            </svg>
                            <span>ลบ</span>
                        </button>
                    </div>

                    <div class="side-description">
                        <h3>รายละเอียดงาน</h3>
                        <p>รายละเอียดเพิ่มเติมเกี่ยวกับตำแหน่งงานนี้จะแสดงที่นี่</p>
                    </div>
                </div>
            </div>
        </div>

        <!-- Delete Confirmation Modal -->
        <div v-if="showDeleteConfirm" class="modal-overlay" @click.self="showDeleteConfirm = false">
            <div class="modal-content delete-confirm-modal">
                <div class="delete-icon">
                    <svg xmlns="http://www.w3.org/2000/svg" width="48" height="48" viewBox="0 0 24 24" fill="none"
                        stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                        <circle cx="12" cy="12" r="10"></circle>
                        <line x1="15" y1="9" x2="9" y2="15"></line>
                        <line x1="9" y1="9" x2="15" y2="15"></line>
                    </svg>
                </div>
                <h3>ยืนยันการลบประกาศงาน</h3>
                <p v-if="jobToDelete">
                    คุณต้องการลบประกาศงาน <strong>"{{ jobToDelete.title }}"</strong> ใช่หรือไม่?
                </p>
                <p v-else>
                    คุณต้องการลบประกาศงานที่เลือก <strong>{{ selectedJobs.length }} รายการ</strong> ใช่หรือไม่?
                </p>
                <p class="warning-text">การดำเนินการนี้ไม่สามารถย้อนกลับได้</p>
                <div class="modal-actions">
                    <button class="btn-cancel" @click="showDeleteConfirm = false">ยกเลิก</button>
                    <button class="btn-confirm-delete" @click="executeDelete">ยืนยันการลบ</button>
                </div>
            </div>
        </div>
    </div>
</template>

<script setup>
import { ref, computed, watch, onMounted, onBeforeUnmount } from 'vue';

// =======================================================
// STATE
// =======================================================
const isLoading = ref(false);
const activeTab = ref('all');
const searchText = ref('');
const selectedYear = ref('ทั้งหมด');
const selectedSemester = ref('ทั้งหมด');
const selectedDepartment = ref('ทั้งหมด');
const sortAscending = ref(false);
const currentPage = ref(1);
const itemsPerPage = ref(10);

// Dropdown states
const isYearFilterOpen = ref(false);
const isSemesterFilterOpen = ref(false);
const isDepartmentFilterOpen = ref(false);
const yearFilterRef = ref(null);
const semesterFilterRef = ref(null);
const departmentFilterRef = ref(null);

// Selection states
const selectAll = ref(false);
const selectedJobs = ref([]);

// Modal states
const selectedJobDetail = ref(null);
const showDeleteConfirm = ref(false);
const jobToDelete = ref(null);

// Data
const academicYears = ['ทั้งหมด', 2568, 2567, 2566, 2565, 2564, 2563, 2562, 2561, 2560, 2559];
const semesters = ['ทั้งหมด', 1, 2, 3];
const departments = [
    'ทั้งหมด',
    'ภาควิชาคณิตศาสตร์',
    'ภาควิชาเคมี',
    'ภาควิชาชีววิทยา',
    'ภาควิชาฟิสิกส์',
    'ภาควิชาวิทยาศาสตร์สิ่งแวดล้อม',
    'ภาควิชาสถิติ',
    'ภาควิชาคอมพิวเตอร์',
    'ภาควิชาจุลชีววิทยา'
];

// Mock data
const mockJobsData = ref([
    { id: 1, year: 2568, semester: 1, title: 'Software Engineer', department: 'ภาควิชาคอมพิวเตอร์', status: 'ผ่านการตรวจสอบ', duration: 30, postedDate: '15/01/2568', applicants: 24 },
    { id: 2, year: 2568, semester: 2, title: 'Data Scientist', department: 'ภาควิชาสถิติ', status: 'รอตรวจสอบ', duration: 15, postedDate: '', applicants: 0 },
    { id: 3, year: 2568, semester: 1, title: 'Research Assistant', department: 'ภาควิชาชีววิทยา', status: 'ผ่านการตรวจสอบ', duration: 45, postedDate: '10/01/2568', applicants: 8 },
    { id: 4, year: 2567, semester: 2, title: 'Lab Technician', department: 'ภาควิชาเคมี', status: 'ไม่ผ่านการตรวจสอบ', duration: 7, postedDate: '18/01/2568', applicants: 0 },
    { id: 5, year: 2567, semester: 2, title: 'Physics Instructor', department: 'ภาควิชาฟิสิกส์', status: 'ผ่านการตรวจสอบ', duration: 60, postedDate: '05/01/2568', applicants: 35 },
    { id: 6, year: 2568, semester: 3, title: 'Environmental Analyst', department: 'ภาควิชาวิทยาศาสตร์สิ่งแวดล้อม', status: 'รอตรวจสอบ', duration: 20, postedDate: '', applicants: 0 },
    { id: 7, year: 2567, semester: 1, title: 'Mathematics Tutor', department: 'ภาควิชาคณิตศาสตร์', status: 'ผ่านการตรวจสอบ', duration: 30, postedDate: '12/01/2568', applicants: 18 },
    { id: 8, year: 2568, semester: 1, title: 'Microbiologist', department: 'ภาควิชาจุลชีววิทยา', status: 'ผ่านการตรวจสอบ', duration: 25, postedDate: '08/01/2568', applicants: 21 },
    { id: 9, year: 2568, semester: 2, title: 'Quality Assurance', department: 'ภาควิชาเคมี', status: 'รอตรวจสอบ', duration: 14, postedDate: '', applicants: 0 },
    { id: 10, year: 2568, semester: 1, title: 'IT Support', department: 'ภาควิชาคอมพิวเตอร์', status: 'ผ่านการตรวจสอบ', duration: 30, postedDate: '16/01/2568', applicants: 42 },
    { id: 11, year: 2568, semester: 3, title: 'Statistician', department: 'ภาควิชาสถิติ', status: 'ไม่ผ่านการตรวจสอบ', duration: 10, postedDate: '28/01/2568', applicants: 3 },
    { id: 12, year: 2568, semester: 3, title: 'Biology Teacher', department: 'ภาควิชาชีววิทยา', status: 'ผ่านการตรวจสอบ', duration: 40, postedDate: '03/01/2568', applicants: 29 },
    { id: 13, year: 2568, semester: 1, title: 'Chemistry Instructor', department: 'ภาควิชาเคมี', status: 'รอตรวจสอบ', duration: 18, postedDate: '', applicants: 0 },
]);

// =======================================================
// COMPUTED PROPERTIES
// =======================================================
const tabs = computed(() => [
    { value: 'all', label: 'ทั้งหมด', count: mockJobsData.value.length },
    { value: 'ผ่านการตรวจสอบ', label: 'ผ่านการตรวจสอบ', count: approvedCount.value },
    { value: 'รอตรวจสอบ', label: 'รอตรวจสอบ', count: pendingCount.value },
    { value: 'ไม่ผ่านการตรวจสอบ', label: 'ไม่ผ่านการตรวจสอบ', count: rejectedCount.value }
]);

const approvedCount = computed(() =>
    mockJobsData.value.filter(j => j.status === 'ผ่านการตรวจสอบ').length
);

const pendingCount = computed(() =>
    mockJobsData.value.filter(j => j.status === 'รอตรวจสอบ').length
);

const rejectedCount = computed(() =>
    mockJobsData.value.filter(j => j.status === 'ไม่ผ่านการตรวจสอบ').length
);

const totalCount = computed(() => mockJobsData.value.length);

const filteredJobs = computed(() => {
    let result = [...mockJobsData.value];

    // Filter by tab
    if (activeTab.value !== 'all') {
        result = result.filter(job => job.status === activeTab.value);
    }

    // Filter by search
    if (searchText.value) {
        const search = searchText.value.toLowerCase();
        result = result.filter(job =>
            job.title.toLowerCase().includes(search)
        );
    }

    // Filter by year
    if (selectedYear.value !== 'ทั้งหมด') {
        result = result.filter(job => job.year === selectedYear.value);
    }

    // Filter by semester
    if (selectedSemester.value !== 'ทั้งหมด') {
        result = result.filter(job => job.semester === selectedSemester.value);
    }

    // Filter by department
    if (selectedDepartment.value !== 'ทั้งหมด') {
        result = result.filter(job => job.department === selectedDepartment.value);
    }

    // Sort
    result.sort((a, b) => {
        const dateA = a.postedDate ? new Date(a.postedDate.split('/').reverse().join('-')) : new Date(0);
        const dateB = b.postedDate ? new Date(b.postedDate.split('/').reverse().join('-')) : new Date(0);
        return sortAscending.value ? dateA - dateB : dateB - dateA;
    });

    return result;
});

const totalPages = computed(() =>
    Math.ceil(filteredJobs.value.length / itemsPerPage.value)
);

const paginatedJobs = computed(() => {
    const start = (currentPage.value - 1) * itemsPerPage.value;
    const end = start + itemsPerPage.value;
    return filteredJobs.value.slice(start, end);
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

const hasActiveFilters = computed(() =>
    selectedYear.value !== 'ทั้งหมด' ||
    selectedSemester.value !== 'ทั้งหมด' ||
    selectedDepartment.value !== 'ทั้งหมด' ||
    searchText.value
);

// =======================================================
// METHODS
// =======================================================
const getStatusClass = (status) => {
    const statusMap = {
        'รอตรวจสอบ': 'pending',
        'ผ่านการตรวจสอบ': 'approved',
        'ไม่ผ่านการตรวจสอบ': 'rejected'
    };
    return statusMap[status] || '';
};

const getEmptyMessage = () => {
    if (searchText.value) {
        return `ไม่พบประกาศงานที่ตรงกับคำค้นหา "${searchText.value}"`;
    }
    if (activeTab.value === 'ผ่านการตรวจสอบ') {
        return 'ยังไม่มีประกาศงานที่ผ่านการตรวจสอบ';
    }
    if (activeTab.value === 'รอตรวจสอบ') {
        return 'ยังไม่มีประกาศงานที่รอตรวจสอบ';
    }
    if (activeTab.value === 'ไม่ผ่านการตรวจสอบ') {
        return 'ยังไม่มีประกาศงานที่ไม่ผ่านการตรวจสอบ';
    }
    return 'ยังไม่มีประกาศงานในระบบ';
};

const toggleSort = () => {
    sortAscending.value = !sortAscending.value;
};

const clearAllFilters = () => {
    searchText.value = '';
    selectedYear.value = 'ทั้งหมด';
    selectedSemester.value = 'ทั้งหมด';
    selectedDepartment.value = 'ทั้งหมด';
};

const handleSearch = () => {
    currentPage.value = 1;
};

const toggleSelectAll = () => {
    if (selectAll.value) {
        selectedJobs.value = paginatedJobs.value.map(job => job.id);
    } else {
        selectedJobs.value = [];
    }
};

const openSidePanel = (job) => {
    selectedJobDetail.value = job;
};

const closeSidePanel = () => {
    selectedJobDetail.value = null;
};

const createNewJob = () => {
    console.log('Create new job');
    // TODO: Navigate to create job page
};

const editJob = (job) => {
    console.log('Edit job:', job);
    // TODO: Navigate to edit job page
};

const confirmDelete = (job) => {
    jobToDelete.value = job;
    showDeleteConfirm.value = true;
    if (selectedJobDetail.value) {
        closeSidePanel();
    }
};

const executeDelete = () => {
    if (jobToDelete.value) {
        // Delete single job
        const index = mockJobsData.value.findIndex(j => j.id === jobToDelete.value.id);
        if (index > -1) {
            mockJobsData.value.splice(index, 1);
        }
    } else {
        // Delete multiple jobs
        mockJobsData.value = mockJobsData.value.filter(
            job => !selectedJobs.value.includes(job.id)
        );
        selectedJobs.value = [];
    }

    showDeleteConfirm.value = false;
    jobToDelete.value = null;
};

const bulkDelete = () => {
    jobToDelete.value = null;
    showDeleteConfirm.value = true;
};

const handleClickOutside = (event) => {
    if (yearFilterRef.value && !yearFilterRef.value.contains(event.target)) {
        isYearFilterOpen.value = false;
    }
    if (semesterFilterRef.value && !semesterFilterRef.value.contains(event.target)) {
        isSemesterFilterOpen.value = false;
    }
    if (departmentFilterRef.value && !departmentFilterRef.value.contains(event.target)) {
        isDepartmentFilterOpen.value = false;
    }
};

// =======================================================
// WATCHERS
// =======================================================
watch([activeTab, searchText, selectedYear, selectedSemester, selectedDepartment], () => {
    currentPage.value = 1;
    selectedJobs.value = [];
    selectAll.value = false;
});

watch(paginatedJobs, () => {
    selectAll.value = false;
    selectedJobs.value = selectedJobs.value.filter(id =>
        paginatedJobs.value.some(job => job.id === id)
    );
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

.manage-jobs-page {
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

.btn-create {
    display: flex;
    align-items: center;
    gap: 8px;
    padding: 12px 24px;
    background: linear-gradient(135deg, #037266 0%, #026b5b 100%);
    color: #fff;
    border: none;
    border-radius: 12px;
    font-size: 15px;
    font-weight: 600;
    cursor: pointer;
    transition: all 0.2s;
    box-shadow: 0 4px 12px rgba(3, 114, 102, 0.2);
    white-space: nowrap;
}

.btn-create:hover {
    transform: translateY(-2px);
    box-shadow: 0 6px 16px rgba(3, 114, 102, 0.3);
}

/* =================================== */
/* STATS OVERVIEW */
/* =================================== */
.stats-overview {
    display: grid;
    grid-template-columns: repeat(auto-fit, minmax(240px, 1fr));
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

.stat-icon.approved {
    background: linear-gradient(135deg, #10b981 0%, #059669 100%);
    color: #fff;
}

.stat-icon.pending {
    background: linear-gradient(135deg, #f59e0b 0%, #d97706 100%);
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

.search-sort {
    display: flex;
    align-items: center;
    margin: 25px 0;
    flex-wrap: wrap;
    gap: 12px;
}

.search-box {
    flex: 1;
    min-width: 250px;
    position: relative;
    display: flex;
    align-items: center;
}

.search-box svg {
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

.filter-dropdown {
    position: relative;
}

.btn-filter-extra {
    display: flex;
    align-items: center;
    gap: 8px;
    padding: 9px 16px;
    background: #f1f1f1;
    border: none;
    border-radius: 10px;
    font-size: 14px;
    font-weight: 600;
    color: #555;
    cursor: pointer;
    transition: all 0.2s;
    white-space: nowrap;
    height: 45px;
    min-width: 150px;
}

.btn-filter-extra:hover {
    background: #e0e0e0;
}

.filter-count {
    background: #037266;
    color: #fff;
    padding: 2px 8px;
    border-radius: 12px;
    font-size: 12px;
    font-weight: 700;
}

.filter-menu {
    position: absolute;
    top: calc(100% + 8px);
    left: 0;
    background: #fff;
    border: 1px solid #e5e7e6;
    border-radius: 12px;
    padding: 12px;
    min-width: 200px;
    max-height: 300px;
    overflow-y: auto;
    box-shadow: 0 8px 24px rgba(6, 20, 18, 0.12);
    z-index: 100;
}

.filter-radio {
    display: flex;
    align-items: center;
    gap: 10px;
    padding: 10px;
    cursor: pointer;
    border-radius: 8px;
    transition: all 0.2s;
}

.filter-radio:hover {
    background: #f7f9f8;
}

.filter-radio input[type="radio"] {
    width: 18px;
    height: 18px;
    cursor: pointer;
    accent-color: #037266;
}

.filter-radio span {
    font-size: 14px;
    color: #333;
    font-weight: 500;
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

.btn-clear-all {
    display: flex;
    align-items: center;
    gap: 6px;
    padding: 9px 16px;
    background: #fee2e2;
    border: 2px solid #fecaca;
    border-radius: 10px;
    font-size: 14px;
    font-weight: 600;
    color: #dc2626;
    cursor: pointer;
    transition: all 0.2s;
}

.btn-clear-all:hover {
    background: #fecaca;
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

.btn-create-empty {
    margin-top: 12px;
    display: flex;
    align-items: center;
    gap: 8px;
    padding: 12px 24px;
    background: linear-gradient(135deg, #037266 0%, #026b5b 100%);
    color: #fff;
    border: none;
    border-radius: 12px;
    font-size: 15px;
    font-weight: 600;
    cursor: pointer;
    transition: all 0.2s;
}

.btn-create-empty:hover {
    transform: translateY(-2px);
    box-shadow: 0 6px 16px rgba(3, 114, 102, 0.3);
}

/* =================================== */
/* JOBS TABLE */
/* =================================== */
.jobs-table-container {
    background: #fff;
    border-radius: 16px;
    overflow: hidden;
    box-shadow: 0 2px 12px rgba(3, 114, 102, 0.07);
    border: 1px solid #eef0ef;
}

.jobs-table {
    width: 100%;
    border-collapse: collapse;
}

.jobs-table thead {
    background: #f7f9f8;
}

.jobs-table th {
    padding: 16px;
    text-align: left;
    font-size: 14px;
    font-weight: 700;
    color: #555;
    text-transform: uppercase;
    letter-spacing: 0.5px;
    border-bottom: 2px solid #e5e5e5;
}

.jobs-table tbody tr {
    border-bottom: 1px solid #f0f0f0;
    transition: all 0.2s;
}

.jobs-table tbody tr:hover {
    background: #f9fafb;
}

.jobs-table tbody tr.selected {
    background: #f0f8f7;
}

.jobs-table td {
    padding: 16px;
    vertical-align: middle;
}

.col-checkbox {
    width: 50px;
    text-align: center;
}

.col-checkbox input[type="checkbox"] {
    width: 18px;
    height: 18px;
    cursor: pointer;
    accent-color: #037266;
}

.col-title {
    min-width: 250px;
}

.title-cell {
    display: flex;
    gap: 14px;
    align-items: flex-start;
}

.title-info {
    flex: 1;
    min-width: 0;
}

.title-info h4 {
    font-size: 15px;
    font-weight: 600;
    color: #222;
    margin: 0 0 6px 0;
    line-height: 1.4;
}

.title-info p {
    font-size: 13px;
    color: #666;
    margin: 0;
    line-height: 1.5;
}

.col-department {
    width: 200px;
}

.department-text {
    font-size: 14px;
    color: #555;
    font-weight: 500;
}

.col-status {
    width: 160px;
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

.status-badge.approved {
    background: #d1fae5;
    color: #065f46;
}

.status-badge.pending {
    background: #fef3c7;
    color: #92400e;
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

.col-duration {
    width: 140px;
}

.duration-cell {
    display: flex;
    align-items: center;
    gap: 6px;
    font-size: 13px;
    color: #666;
}

.duration-cell svg {
    flex-shrink: 0;
    color: #999;
}

.col-date {
    width: 140px;
}

.date-cell {
    display: flex;
    align-items: center;
    gap: 6px;
    font-size: 13px;
    color: #666;
}

.date-cell svg {
    flex-shrink: 0;
    color: #999;
}

.col-applicants {
    width: 120px;
}

.applicants-cell {
    display: flex;
    align-items: center;
    gap: 6px;
    font-size: 13px;
    color: #666;
    font-weight: 600;
}

.applicants-cell svg {
    flex-shrink: 0;
    color: #999;
}

.col-actions {
    width: 140px;
}

.action-buttons {
    display: flex;
    gap: 6px;
}

.btn-action {
    width: 36px;
    height: 36px;
    border: none;
    border-radius: 8px;
    display: flex;
    align-items: center;
    justify-content: center;
    cursor: pointer;
    transition: all 0.2s;
}

.btn-action.btn-view {
    background: #e0f2fe;
    color: #0369a1;
}

.btn-action.btn-view:hover {
    background: #bae6fd;
    transform: scale(1.05);
}

.btn-action.btn-edit {
    background: #fef3c7;
    color: #d97706;
}

.btn-action.btn-edit:hover {
    background: #fde68a;
    transform: scale(1.05);
}

.btn-action.btn-delete {
    background: #fee2e2;
    color: #dc2626;
}

.btn-action.btn-delete:hover {
    background: #fecaca;
    transform: scale(1.05);
}

/* =================================== */
/* BULK ACTIONS BAR */
/* =================================== */
.bulk-actions-bar {
    position: fixed;
    bottom: -100px;
    left: 50%;
    transform: translateX(-50%);
    background: #fff;
    border-radius: 16px;
    padding: 16px 24px;
    box-shadow: 0 8px 32px rgba(0, 0, 0, 0.15);
    border: 1px solid #e5e5e5;
    display: flex;
    align-items: center;
    gap: 20px;
    transition: bottom 0.3s ease;
    z-index: 1000;
    max-width: 90%;
}

.bulk-actions-bar.visible {
    bottom: 30px;
}

.bulk-info {
    font-size: 15px;
    font-weight: 600;
    color: #333;
    white-space: nowrap;
}

.bulk-buttons {
    display: flex;
    gap: 10px;
}

.btn-bulk {
    display: flex;
    align-items: center;
    gap: 6px;
    padding: 10px 18px;
    border: none;
    border-radius: 10px;
    font-size: 14px;
    font-weight: 600;
    cursor: pointer;
    transition: all 0.2s;
    white-space: nowrap;
}

.btn-bulk.btn-delete-bulk {
    background: #fee2e2;
    color: #dc2626;
}

.btn-bulk.btn-delete-bulk:hover {
    background: #fecaca;
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
/* SIDE PANEL */
/* =================================== */
.side-panel-overlay {
    position: fixed;
    inset: 0;
    background: rgba(0, 0, 0, 0.4);
    display: flex;
    justify-content: flex-end;
    z-index: 2000;
    animation: fadeIn 0.2s ease;
}

@keyframes fadeIn {
    from {
        opacity: 0;
    }
    to {
        opacity: 1;
    }
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
    z-index: 10;
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
    margin-bottom: 12px;
    color: #222;
    padding-right: 40px;
    line-height: 1.3;
}

.job-meta-inline {
    display: flex;
    gap: 12px;
    margin-bottom: 12px;
    align-items: center;
    flex-wrap: wrap;
}

.department-text-meta {
    font-size: 13px;
    color: #666;
    font-weight: 500;
}

.job-details-grid {
    display: grid;
    grid-template-columns: repeat(2, 1fr);
    gap: 12px;
    margin-top: 16px;
}

.detail-item {
    display: flex;
    align-items: center;
    gap: 8px;
    font-size: 13px;
    color: #555;
}

.detail-item svg {
    flex-shrink: 0;
    color: #999;
}

.side-action-buttons {
    display: flex;
    gap: 12px;
    margin: 20px 0 10px 0;
    flex-wrap: wrap;
}

.btn-side-action {
    flex: 1;
    min-width: 140px;
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
}

.btn-side-action.btn-edit {
    background: #ffffff;
    color: #d97706;
    border: 2px solid #fde68a;
}

.btn-side-action.btn-edit:hover {
    background: #fef3c7;
    border-color: #d97706;
    transform: translateY(-2px);
    box-shadow: 0 6px 20px rgba(217, 119, 6, 0.15);
}

.btn-side-action.btn-delete {
    background: #ffffff;
    color: #dc2626;
    border: 2px solid #fecaca;
}

.btn-side-action.btn-delete:hover {
    background: #fee2e2;
    border-color: #dc2626;
    transform: translateY(-2px);
    box-shadow: 0 6px 20px rgba(220, 38, 38, 0.15);
}

.side-description {
    margin-top: 15px;
}

.side-description h3 {
    font-size: 18px;
    font-weight: 600;
    color: #222;
    margin-bottom: 12px;
}

.side-description p {
    font-size: 15px;
    line-height: 1.7;
    color: #444;
    white-space: pre-line;
}

/* =================================== */
/* DELETE CONFIRMATION MODAL */
/* =================================== */
.modal-overlay {
    position: fixed;
    inset: 0;
    background: rgba(0, 0, 0, 0.5);
    display: flex;
    align-items: center;
    justify-content: center;
    z-index: 2500;
    padding: 20px;
    animation: fadeIn 0.2s ease;
}

.modal-content {
    background: #fff;
    border-radius: 20px;
    max-width: 500px;
    width: 100%;
    position: relative;
    animation: slideUp 0.3s ease;
}

@keyframes slideUp {
    from {
        opacity: 0;
        transform: translateY(20px);
    }
    to {
        opacity: 1;
        transform: translateY(0);
    }
}

.delete-confirm-modal {
    padding: 40px;
    text-align: center;
}

.delete-icon {
    width: 80px;
    height: 80px;
    border-radius: 50%;
    background: #fee2e2;
    color: #dc2626;
    display: flex;
    align-items: center;
    justify-content: center;
    margin: 0 auto 24px;
}

.delete-confirm-modal h3 {
    font-size: 24px;
    font-weight: 700;
    color: #222;
    margin: 0 0 16px 0;
}

.delete-confirm-modal p {
    font-size: 15px;
    color: #666;
    margin: 0 0 12px 0;
    line-height: 1.6;
}

.warning-text {
    color: #dc2626;
    font-weight: 600;
    margin-bottom: 24px;
}

.modal-actions {
    display: flex;
    gap: 12px;
    justify-content: center;
    margin-top: 24px;
}

.btn-cancel,
.btn-confirm-delete {
    flex: 1;
    padding: 12px 24px;
    border: none;
    border-radius: 12px;
    font-size: 15px;
    font-weight: 600;
    cursor: pointer;
    transition: all 0.2s;
}

.btn-cancel {
    background: #f0f0f0;
    color: #666;
}

.btn-cancel:hover {
    background: #e0e0e0;
}

.btn-confirm-delete {
    background: #dc2626;
    color: #fff;
}

.btn-confirm-delete:hover {
    background: #b91c1c;
}

/* =================================== */
/* DROPDOWN TRANSITIONS */
/* =================================== */
.dropdown-fade-enter-active,
.dropdown-fade-leave-active {
    transition: opacity 0.2s ease, transform 0.2s ease;
}

.dropdown-fade-enter-from {
    opacity: 0;
    transform: translateY(-8px);
}

.dropdown-fade-leave-to {
    opacity: 0;
    transform: translateY(-4px);
}

/* =================================== */
/* RESPONSIVE */
/* =================================== */
@media (max-width: 1024px) {
    .jobs-table {
        font-size: 14px;
    }

    .col-title {
        min-width: 200px;
    }
}

@media (max-width: 768px) {
    .manage-jobs-page {
        padding: 15px;
    }

    .page-header {
        flex-direction: column;
        align-items: stretch;
    }

    .btn-create {
        width: 100%;
        justify-content: center;
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

    .btn-filter-extra,
    .btn-sort,
    .btn-clear-all,
    .btn-search {
        width: 100%;
        justify-content: center;
    }

    .jobs-table-container {
        overflow-x: auto;
    }

    .jobs-table {
        min-width: 900px;
    }

    .bulk-actions-bar {
        flex-direction: column;
        align-items: stretch;
        max-width: calc(100% - 30px);
    }

    .bulk-buttons {
        flex-direction: column;
    }

    .btn-bulk {
        width: 100%;
        justify-content: center;
    }

    .job-details-grid {
        grid-template-columns: 1fr;
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
    }

    .tab-btn {
        padding: 8px 12px;
        font-size: 13px;
    }

    .pagination {
        gap: 4px;
    }

    .btn-page {
        min-width: 36px;
        height: 36px;
        font-size: 13px;
    }
}
</style>