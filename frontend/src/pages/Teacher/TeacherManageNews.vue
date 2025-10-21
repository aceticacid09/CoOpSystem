<!-- frontend/src/pages/Teacher/TeacherManageNews.vue -->
<template>
    <DashboardLayout role="teacher">
        <div class="manage-news-page">
            <div class="page-header">
                <div class="header-content">
                    <h1>จัดการข่าวสารและกิจกรรม</h1>
                    <p class="page-subtitle">จัดการ แก้ไข และติดตามสถานะของข่าวสารและกิจกรรมทั้งหมด</p>
                </div>
                <button class="btn-create" @click="router.push({ name: 'TeacherCreateNews' })">
                    <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none"
                        stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                        <line x1="12" y1="5" x2="12" y2="19"></line>
                        <line x1="5" y1="12" x2="19" y2="12"></line>
                    </svg>
                    สร้างข่าวสารใหม่
                </button>
            </div>

            <div class="stats-overview">
                <div class="stat-card">
                    <div class="stat-icon published">
                        <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none"
                            stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                            <path d="M22 11.08V12a10 10 0 1 1-5.93-9.14"></path>
                            <polyline points="22 4 12 14.01 9 11.01"></polyline>
                        </svg>
                    </div>
                    <div class="stat-info">
                        <span class="stat-label">เผยแพร่แล้ว</span>
                        <span class="stat-value">{{ publishedCount }}</span>
                    </div>
                </div>

                <div class="stat-card">
                    <div class="stat-icon draft">
                        <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none"
                            stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                            <path d="M14 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8z"></path>
                            <polyline points="14 2 14 8 20 8"></polyline>
                        </svg>
                    </div>
                    <div class="stat-info">
                        <span class="stat-label">แบบร่าง</span>
                        <span class="stat-value">{{ draftCount }}</span>
                    </div>
                </div>

                <div class="stat-card">
                    <div class="stat-icon scheduled">
                        <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none"
                            stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                            <circle cx="12" cy="12" r="10"></circle>
                            <polyline points="12 6 12 12 16 14"></polyline>
                        </svg>
                    </div>
                    <div class="stat-info">
                        <span class="stat-label">รอเผยแพร่</span>
                        <span class="stat-value">{{ scheduledCount }}</span>
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

                <div class="status-date-filters">
                    <CustomDropdown label="หมวดหมู่" v-model="activeTab" :options="tabs" width="300px" />
                    <CustomDropdown label="สถานะ" v-model="statusFilter" :options="statusOptions" width="220px" />
                    <div class="filter-group">
                        <label class="filter-label">วันที่</label>
                        <div class="date-filter-dropdown" ref="dateFilterRef">
                            <button class="date-dropdown-button" @click.stop="isDateFilterOpen = !isDateFilterOpen">
                                <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24"
                                    fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round"
                                    stroke-linejoin="round">
                                    <rect x="3" y="4" width="18" height="18" rx="2" ry="2" />
                                    <line x1="16" y1="2" x2="16" y2="6" />
                                    <line x1="8" y1="2" x2="8" y2="6" />
                                    <line x1="3" y1="10" x2="21" y2="10" />
                                </svg>
                                <span>{{ dateRangeDisplay }}</span>
                                <svg class="dropdown-arrow" :class="{ open: isDateFilterOpen }"
                                    xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24"
                                    fill="none" stroke="#9CA3AF" stroke-width="2" stroke-linecap="round"
                                    stroke-linejoin="round">
                                    <polyline points="6 9 12 15 18 9"></polyline>
                                </svg>
                            </button>
                            <transition name="dropdown-fade">
                                <div v-if="isDateFilterOpen" class="date-dropdown-menu">
                                    <div class="date-inputs">
                                        <div class="input-group">
                                            <label>จาก</label>
                                            <input type="date" v-model="dateFrom" class="date-input" />
                                        </div>
                                        <div class="input-group">
                                            <label>ถึง</label>
                                            <input type="date" v-model="dateTo" class="date-input" />
                                        </div>
                                    </div>
                                    <div class="date-actions">
                                        <button class="btn-clear-date" @click="clearDateFilter">ล้าง</button>
                                        <button class="btn-apply-date" @click="applyDateFilter">ใช้งาน</button>
                                    </div>
                                </div>
                            </transition>
                        </div>
                    </div>
                </div>

                <SearchBar v-model="searchText" placeholder="ค้นหาหัวข้อข่าว..." :is-ascending="sortAscending"
                    @toggle-sort="toggleSort" />

            </div>

            <div v-if="isLoading" class="loading-state">
                <div class="spinner"></div>
                <p>กำลังโหลดข้อมูล...</p>
            </div>

            <div v-else-if="filteredNews.length === 0" class="empty-state">
                <svg xmlns="http://www.w3.org/2000/svg" width="64" height="64" viewBox="0 0 24 24" fill="none"
                    stroke="currentColor" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round">
                    <path d="M14 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8z"></path>
                    <polyline points="14 2 14 8 20 8"></polyline>
                    <line x1="16" y1="13" x2="8" y2="13"></line>
                    <line x1="16" y1="17" x2="8" y2="17"></line>
                    <polyline points="10 9 9 9 8 9"></polyline>
                </svg>
                <h3>ไม่พบข่าวสาร</h3>
                <p>{{ getEmptyMessage() }}</p>
                <button class="btn-create-empty" @click="router.push({ name: 'TeacherCreateNews' })">
                    <svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none"
                        stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                        <line x1="12" y1="5" x2="12" y2="19"></line>
                        <line x1="5" y1="12" x2="19" y2="12"></line>
                    </svg>
                    สร้างข่าวสารใหม่
                </button>
            </div>

            <div v-else class="news-table-container">
                <table class="news-table">
                    <thead>
                        <tr>
                            <th class="col-checkbox">
                                <input type="checkbox" v-model="selectAll" @change="toggleSelectAll" />
                            </th>
                            <th class="col-title">หัวข้อข่าว</th>
                            <th class="col-category">หมวดหมู่</th>
                            <th class="col-status">สถานะ</th>
                            <th class="col-date">วันที่ประกาศ</th>
                            <th class="col-actions">จัดการ</th>
                        </tr>
                    </thead>
                    <tbody>
                        <tr v-for="news in paginatedNews" :key="news.id" class="news-row"
                            :class="{ selected: selectedNews.includes(news.id) }">
                            <td class="col-checkbox">
                                <input type="checkbox" :value="news.id" v-model="selectedNews" />
                            </td>
                            <td class="col-title">
                                <div class="title-cell">
                                    <div v-if="news.images && news.images.length > 0" class="news-thumbnail">
                                        <img :src="news.images[0]" :alt="news.title" />
                                    </div>
                                    <div class="title-info">
                                        <h4>{{ news.title }}</h4>
                                        <p>{{ truncateText(news.description, 80) }}</p>
                                    </div>
                                </div>
                            </td>
                            <td class="col-category">
                                <span class="category-text">{{ news.category }}</span>
                            </td>
                            <td class="col-status">
                                <span class="status-badge" :class="news.status">
                                    <span class="status-dot"></span>
                                    {{ getStatusText(news.status) }}
                                </span>
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
                                    {{ formatDate(news.publishDate || news.date) }}
                                </div>
                            </td>
                            <td class="col-actions">
                                <div class="action-buttons">
                                    <button class="btn-action btn-view" @click="openSidePanel(news)"
                                        title="ดูรายละเอียด">
                                        <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16"
                                            viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"
                                            stroke-linecap="round" stroke-linejoin="round">
                                            <path d="M1 12s4-8 11-8 11 8 11 8-4 8-11 8-11-8-11-8z"></path>
                                            <circle cx="12" cy="12" r="3"></circle>
                                        </svg>
                                    </button>
                                    <button class="btn-action btn-edit" @click="editNews(news)" title="แก้ไข">
                                        <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16"
                                            viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"
                                            stroke-linecap="round" stroke-linejoin="round">
                                            <path d="M11 4H4a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2v-7">
                                            </path>
                                            <path d="M18.5 2.5a2.121 2.121 0 0 1 3 3L12 15l-4 1 1-4 9.5-9.5z">
                                            </path>
                                        </svg>
                                    </button>
                                    <button class="btn-action btn-delete" @click="confirmDelete(news)" title="ลบ">
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

            <div v-if="filteredNews.length > 0" class="bulk-actions-bar" :class="{ visible: selectedNews.length > 0 }">
                <div class="bulk-info">
                    <span>เลือกแล้ว {{ selectedNews.length }} รายการ</span>
                </div>
                <div class="bulk-buttons">
                    <button class="btn-bulk btn-publish" @click="bulkPublish" v-if="canBulkPublish">
                        <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none"
                            stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                            <path d="M22 11.08V12a10 10 0 1 1-5.93-9.14"></path>
                            <polyline points="22 4 12 14.01 9 11.01"></polyline>
                        </svg>
                        เผยแพร่ทันที
                    </button>
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

            <div v-if="filteredNews.length > itemsPerPage" class="pagination">
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

            <NewsSidePanel v-if="selectedNewsDetail" :news="selectedNewsDetail" :show-actions="false"
                :show-status="true" @close="closeSidePanel">
                <template #actions>
                    <div class="side-action-buttons">
                        <button class="btn-side-action btn-edit" @click="editNews(selectedNewsDetail)">
                            <svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24"
                                fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round"
                                stroke-linejoin="round">
                                <path d="M11 4H4a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2v-7"></path>
                                <path d="M18.5 2.5a2.121 2.121 0 0 1 3 3L12 15l-4 1 1-4 9.5-9.5z"></path>
                            </svg>
                            <span>แก้ไข</span>
                        </button>
                        <button class="btn-side-action btn-delete" @click="confirmDelete(selectedNewsDetail)">
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
                </template>
            </NewsSidePanel>

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
                    <h3>ยืนยันการลบข่าวสาร</h3>
                    <p v-if="newsToDelete">
                        คุณต้องการลบข่าวสาร <strong>"{{ newsToDelete.title }}"</strong> ใช่หรือไม่?
                    </p>
                    <p v-else>
                        คุณต้องการลบข่าวสารที่เลือก <strong>{{ selectedNews.length }} รายการ</strong> ใช่หรือไม่?
                    </p>
                    <p class="warning-text">การดำเนินการนี้ไม่สามารถย้อนกลับได้</p>
                    <div class="modal-actions">
                        <button class="btn-cancel" @click="showDeleteConfirm = false">ยกเลิก</button>
                        <button class="btn-confirm-delete" @click="executeDelete">ยืนยันการลบ</button>
                    </div>
                </div>
            </div>
        </div>
    </DashboardLayout>
</template>

<script setup>
import { ref, computed, watch, onMounted, onBeforeUnmount } from 'vue';
import { useRouter } from 'vue-router';
import DashboardLayout from '../../components/DashboardLayout.vue';
import SearchBar from '../../components/SearchBar.vue';
import NewsSidePanel from '../../components/NewSidePanel.vue';
import CustomDropdown from '../../components/CustomDropdown.vue';
import { API_CONFIG } from '../../../config/api';

const router = useRouter();

// =====================================
// 1. STATE & MOCK DATA
// =====================================
const isLoading = ref(false);
const activeTab = ref('ข่าวทั้งหมด');
const searchText = ref('');
const statusFilter = ref('ทั้งหมด');
const dateFrom = ref('');
const dateTo = ref('');
const sortAscending = ref(false);
const currentPage = ref(1);
const itemsPerPage = ref(10);

// Dropdown/Refs
const isDateFilterOpen = ref(false);
const dateFilterRef = ref(null);

// Selection/Bulk Actions
const selectAll = ref(false);
const selectedNews = ref([]);

// Modal/Side Panel
const selectedNewsDetail = ref(null);
const showDeleteConfirm = ref(false);
const newsToDelete = ref(null);

// Data
// const mockNewsData = ref(getPublishableNews());
const announcements = ref([]);
const fetchAnnouncements = async () => {
  isLoading.value = true;
  try {
    const res = await fetch(`${API_CONFIG.baseURL}${API_CONFIG.endpoints.announcements}`);
    if (!res.ok) throw new Error('Failed to fetch announcements');
    const data = await res.json();
    announcements.value = data.announcements.map(announcement => ({
      id: announcement.post_id,
      post_id: announcement.post_id, // Add this line
      title: announcement.title,
      description: announcement.description,
      status: announcement.status,
      date: announcement.created_at,
      publishDate: announcement.publish_date,
      category: "ประชาสัมพันธ์", // Map category if needed
      teacher: announcement.teacher,
      images: announcement.attachments?.map(att => `/uploads/news/${att.filename}`) || []
    }));
  } catch (error) {
    console.error('Error fetching announcements:', error);
  } finally {
    isLoading.value = false;
  }
};

const tabs = ["ข่าวทั้งหมด", "ประชาสัมพันธ์", "ข่าวด่วน", "กิจกรรม", "ประกาศผลการคัดเลือก"];
const statusOptions = ["ทั้งหมด", "เผยแพร่แล้ว", "รอเผยแพร่", "แบบร่าง"];

let statusCheckInterval;

// =====================================
// 2. COMPUTED PROPERTIES
// =====================================

// Date range display
const dateRangeDisplay = computed(() => {
    if (dateFrom.value && dateTo.value) {
        return `${formatDateShort(dateFrom.value)} - ${formatDateShort(dateTo.value)}`;
    }
    return 'เลือกช่วงวันที่';
});

// Stats
const publishedCount = computed(() =>
    announcements.value.filter(n => n.status === 'immediate').length
);

const scheduledCount = computed(() =>
    announcements.value.filter(n => n.status === 'scheduled').length
);

const draftCount = computed(() =>
    announcements.value.filter(n => n.status === 'draft').length
);

const totalCount = computed(() => announcements.value.length);

// News Filtering & Sorting
const filteredNews = computed(() => {
    let result = [...announcements.value];

    // Filter by tab (category)
    if (activeTab.value !== 'ข่าวทั้งหมด') {
        result = result.filter(news => news.category === activeTab.value);
    }

    // Filter by status from dropdown
    const statusMap = {
        'ทั้งหมด': 'all',
        'เผยแพร่แล้ว': 'immediate',
        'รอเผยแพร่': 'scheduled',
        'แบบร่าง': 'draft',
    };
    const mappedStatus = statusMap[statusFilter.value];

    if (mappedStatus !== 'all') {
        result = result.filter(news => news.status === mappedStatus);
    }

    // Filter by search
    if (searchText.value) {
        const search = searchText.value.toLowerCase();
        result = result.filter(news =>
            news.title.toLowerCase().includes(search) ||
            news.description.toLowerCase().includes(search)
        );
    }

    // Filter by date range
    if (dateFrom.value && dateTo.value) {
        result = result.filter(news => {
            const newsDate = new Date(news.publishDate || news.date);
            const from = new Date(dateFrom.value);
            const to = new Date(dateTo.value);
            to.setHours(23, 59, 59, 999);
            return newsDate >= from && newsDate <= to;
        });
    }

    // Sort
    result.sort((a, b) => {
        const dateA = new Date(a.publishDate || a.date);
        const dateB = new Date(b.publishDate || b.date);
        return sortAscending.value ? dateA - dateB : dateB - dateA;
    });

    return result;
});

// Pagination
const totalPages = computed(() =>
    Math.ceil(filteredNews.value.length / itemsPerPage.value)
);

const paginatedNews = computed(() => {
    const start = (currentPage.value - 1) * itemsPerPage.value;
    const end = start + itemsPerPage.value;
    return filteredNews.value.slice(start, end);
});

const visiblePages = computed(() => {
    const pages = [];
    const total = totalPages.value;
    const current = currentPage.value;
    const maxVisible = 5;

    if (total <= maxVisible) {
        for (let i = 1; i <= total; i++) pages.push(i);
        return pages;
    }

    pages.push(1);

    if (current > 3) pages.push('...');

    let start = Math.max(2, current - 1);
    let end = Math.min(total - 1, current + 1);

    if (current <= 3) end = 4;
    if (current >= total - 2) start = total - 3;

    for (let i = start; i <= end; i++) {
        if (!pages.includes(i)) pages.push(i);
    }

    if (current < total - 2) pages.push('...');

    if (total > 1 && !pages.includes(total)) pages.push(total);

    return pages.filter((page, index, self) =>
        (typeof page === 'number' || (page === '...' && self[index - 1] !== '...'))
    );
});

const hasActiveFilters = computed(() =>
    statusFilter.value !== 'ทั้งหมด' ||
    dateFrom.value ||
    dateTo.value ||
    searchText.value
);

const canBulkPublish = computed(() => {
    if (selectedNews.value.length === 0) return false;
    return selectedNews.value.some(id => {
        const news = announcements.value.find(n => n.id === id);
        return news && news.status === 'scheduled';
    });
});

// =====================================
// 3. METHODS & FUNCTIONS
// =====================================

const updateScheduledNews = () => {
    const today = new Date();
    today.setHours(0, 0, 0, 0);

    let updated = false;

    announcements.value = announcements.value.map(news => {
        const newsDate = new Date(news.publishDate || news.date);
        newsDate.setHours(0, 0, 0, 0);

        if (news.status === 'scheduled' && newsDate <= today) {
            updated = true;
            return { ...news, status: 'published' };
        }
        return news;
    });

    if (updated) {
        announcements.value = [...announcements.value];
    }
};

const formatDate = (dateStr) => {
    const d = new Date(dateStr);
    if (isNaN(d)) return '';
    const day = d.getDate();
    const months = [
        "ม.ค.", "ก.พ.", "มี.ค.", "เม.ย.", "พ.ค.", "มิ.ย.",
        "ก.ค.", "ส.ค.", "ก.ย.", "ต.ค.", "พ.ย.", "ธ.ค."
    ];
    const month = months[d.getMonth()];
    const year = d.getFullYear() + 543;
    return `${day} ${month} ${year}`;
};

const formatDateShort = (dateStr) => {
    const d = new Date(dateStr);
    if (isNaN(d)) return '';
    const day = d.getDate();
    const month = d.getMonth() + 1;
    const year = d.getFullYear() + 543;
    return `${day}/${month}/${year}`;
};

const truncateText = (text, maxLength) => {
    if (!text) return '';
    const plainText = text.replace(/<[^>]*>/g, '');
    if (plainText.length <= maxLength) return plainText;
    return plainText.substring(0, maxLength).trim() + '...';
};

const getStatusText = (status) => {
    const statusMap = {
        immediate: 'เผยแพร่แล้ว',
        scheduled: 'รอเผยแพร่',
        draft: 'แบบร่าง'
    };
    return statusMap[status] || status;
};

const getEmptyMessage = () => {
    if (searchText.value) {
        return `ไม่พบข่าวสารที่ตรงกับคำค้นหา "${searchText.value}"`;
    }
    if (activeTab.value !== 'ข่าวทั้งหมด') {
        return `ไม่พบข่าวสารในหมวดหมู่ "${activeTab.value}"`;
    }
    if (statusFilter.value === 'เผยแพร่แล้ว') {
        return 'ยังไม่มีข่าวสารที่เผยแพร่';
    }
    if (statusFilter.value === 'รอเผยแพร่') {
        return 'ยังไม่มีข่าวสารที่รอเผยแพร่';
    }
    return 'ยังไม่มีข่าวสารในระบบ';
};

const toggleSort = () => {
    sortAscending.value = !sortAscending.value;
};

const applyDateFilter = () => {
    isDateFilterOpen.value = false;
};

const clearDateFilter = () => {
    dateFrom.value = '';
    dateTo.value = '';
};

const clearAllFilters = () => {
    searchText.value = '';
    statusFilter.value = 'ทั้งหมด';
    clearDateFilter();
};

const toggleSelectAll = () => {
    if (selectAll.value) {
        selectedNews.value = paginatedNews.value.map(news => news.id);
    } else {
        selectedNews.value = [];
    }
};


const openSidePanel = (news) => {
    selectedNewsDetail.value = news;
};

const closeSidePanel = () => {
    selectedNewsDetail.value = null;
};

const editNews = async (news) => {
  try {
    // First fetch the full announcement details
    const response = await fetch(`${API_CONFIG.baseURL}${API_CONFIG.endpoints.announcements}/${news.id}`);
    if (!response.ok) throw new Error('Failed to fetch announcement details');
    
    const fullNews = await response.json();
    console.log('Fetched announcement details:', fullNews);
    
    // If it's a draft, go to Create with draft data
    if (news.status === 'draft') {
      // Instead of using btoa, use encodeURIComponent
      router.push({
        name: 'TeacherCreateNews',
        query: { 
          draftId: news.post_id,
          data: encodeURIComponent(JSON.stringify(fullNews))
        }
      });
    } else {
      // For published/scheduled news, go to Edit route
      router.push({ 
        name: 'TeacherEditNews', 
        params: { id: news.post_id },
        // Pass data through router state instead of query
        state: { announcement: fullNews }
      });
    }
    closeSidePanel();
  } catch (error) {
    console.error('Error preparing to edit announcement:', error);
    alert('ไม่สามารถแก้ไขข่าวสารได้ กรุณาลองใหม่อีกครั้ง');
  }
};

const confirmDelete = (news) => {
    newsToDelete.value = news;
    showDeleteConfirm.value = true;
    if (selectedNewsDetail.value) {
        closeSidePanel();
    }
};

const executeDelete = async () => {
  try {
    if (newsToDelete.value) {
      const response = await fetch(`${API_CONFIG.baseURL}${API_CONFIG.endpoints.announcements}/${newsToDelete.value.id}`, {
        method: 'DELETE'
      });
      
      if (!response.ok) throw new Error('Failed to delete announcement');
      
      await fetchAnnouncements(); // Refresh the list
    } else if (selectedNews.value.length > 0) {
      // Bulk delete
      await Promise.all(selectedNews.value.map(id => 
        fetch(`${API_CONFIG.baseURL}${API_CONFIG.endpoints.announcements}/${id}`, {
          method: 'DELETE'
        })
      ));
      
      await fetchAnnouncements(); // Refresh the list
      selectedNews.value = [];
      selectAll.value = false;
    }

    showDeleteConfirm.value = false;
    newsToDelete.value = null;
  } catch (error) {
    console.error('Error deleting announcement(s):', error);
    // Add error handling UI feedback here
  }
};

const bulkPublish = async () => {
  try {
    await Promise.all(selectedNews.value.map(async (id) => {
      const news = announcements.value.find(n => n.id === id);
      if (news && news.status === 'scheduled') {
        const response = await fetch(`${API_CONFIG.baseURL}${API_CONFIG.endpoints.announcements}/${id}`, {
          method: 'PUT',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify({
            status: 'published',
            publish_date: new Date().toISOString()
          })
        });
        
        if (!response.ok) throw new Error(`Failed to publish announcement ${id}`);
      }
    }));

    await fetchAnnouncements(); // Refresh the list
    selectedNews.value = [];
    selectAll.value = false;
  } catch (error) {
    console.error('Error publishing announcements:', error);
    // Add error handling UI feedback here
  }
};

const bulkDelete = () => {
    newsToDelete.value = null;
    showDeleteConfirm.value = true;
};

const handleClickOutside = (event) => {
    if (isDateFilterOpen.value && dateFilterRef.value && !dateFilterRef.value.contains(event.target)) {
        isDateFilterOpen.value = false;
    }
};

// =====================================
// 4. WATCHERS & LIFECYCLE HOOKS
// =====================================

watch([activeTab, searchText, statusFilter, dateFrom, dateTo], () => {
    currentPage.value = 1;
    selectedNews.value = [];
    selectAll.value = false;
});

watch(paginatedNews, () => {
    selectAll.value = false;
    selectedNews.value = selectedNews.value.filter(id =>
        paginatedNews.value.some(news => news.id === id)
    );
});

// In the same file
onMounted(async () => {
  document.addEventListener('click', handleClickOutside);
  await fetchAnnouncements();
  
  // Update scheduled announcements periodically
  updateScheduledNews();
  statusCheckInterval = setInterval(updateScheduledNews, 60000);
});

onBeforeUnmount(() => {
    document.removeEventListener('click', handleClickOutside);
    if (statusCheckInterval) {
        clearInterval(statusCheckInterval);
    }
});
</script>

<style scoped>
@import url("https://fonts.googleapis.com/css2?family=Kanit:wght@400;500;600;700&display=swap");

* {
    font-family: "Kanit", sans-serif;
    box-sizing: border-box;
}

.manage-news-page {
    padding: 20px;
    max-width: 1400px;
    margin: 0 auto;
}

/* =================================== */
/* 1. PAGE HEADER */
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
/* 2. STATS OVERVIEW */
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

.stat-icon.published {
    background: linear-gradient(135deg, #10b981 0%, #059669 100%);
    color: #fff;
}

.stat-icon.scheduled {
    background: linear-gradient(135deg, #f59e0b 0%, #d97706 100%);
    color: #fff;
}

.stat-icon.total {
    background: linear-gradient(135deg, #037266 0%, #026b5b 100%);
    color: #fff;
}

.stat-icon.draft {
    background: linear-gradient(135deg, #6b7280 0%, #4b5563 100%);
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
/* 3. FILTER SECTION */
/* =================================== */
.filter-section {
    background: #fff;
    border-radius: 16px;
    padding: 20px;
    margin-bottom: 25px;
    box-shadow: 0 2px 12px rgba(3, 114, 102, 0.07);
    border: 1px solid #eef0ef;
}

.status-date-filters {
    display: flex;
    gap: 20px;
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

.date-filter-dropdown {
    position: relative;
    width: 400px;
    margin-bottom: 10px;
}

.date-dropdown-button {
    display: flex;
    align-items: center;
    justify-content: space-between;
    gap: 10px;
    width: 100%;
    padding: 9px 14px;
    background: #fff;
    border: 1px solid #e5e7e6;
    border-radius: 10px;
    font-size: 14px;
    color: #333;
    font-weight: 500;
    cursor: pointer;
    transition: all 0.2s ease;
    box-shadow: 0 1px 3px rgba(10, 10, 10, 0.02);
}

.date-dropdown-button:hover {
    border-color: #037266;
    box-shadow: 0 4px 12px rgba(3, 114, 102, 0.08);
}

.date-dropdown-button svg:first-child {
    flex-shrink: 0;
    color: #666;
}

.date-dropdown-button span {
    flex: 1;
    text-align: left;
    white-space: nowrap;
    overflow: hidden;
    text-overflow: ellipsis;
}

.dropdown-arrow {
    transition: transform 0.25s ease;
    flex-shrink: 0;
    margin-left: 8px;
}

.dropdown-arrow.open {
    transform: rotate(180deg);
}

.date-dropdown-menu {
    position: absolute;
    top: calc(100% + 6px);
    left: 0;
    width: 100%;
    background: #fff;
    border: 1px solid #e5e7e6;
    border-radius: 10px;
    box-shadow: 0 8px 24px rgba(6, 20, 18, 0.12);
    z-index: 100;
    padding: 12px;
}

.date-inputs {
    display: flex;
    gap: 12px;
    margin-bottom: 12px;
}

.input-group {
    display: flex;
    flex-direction: column;
    gap: 6px;
    flex: 1;
}

.input-group label {
    font-size: 13px;
    color: #666;
    font-weight: 600;
}

.date-input {
    padding: 10px 12px;
    border: 2px solid #e5e5e5;
    border-radius: 8px;
    font-size: 14px;
    font-family: "Kanit", sans-serif;
    transition: all 0.2s;
}

.date-input:focus {
    outline: none;
    border-color: #037266;
}

.date-actions {
    display: flex;
    gap: 8px;
    padding-top: 12px;
    border-top: 1px solid #f0f0f0;
}

.btn-clear-date,
.btn-apply-date {
    flex: 1;
    padding: 8px 16px;
    border: none;
    border-radius: 8px;
    font-size: 14px;
    font-weight: 600;
    cursor: pointer;
    transition: all 0.2s;
}

.btn-clear-date {
    background: #f0f0f0;
    color: #666;
}

.btn-clear-date:hover {
    background: #e0e0e0;
}

.btn-apply-date {
    background: #037266;
    color: #fff;
}

.btn-apply-date:hover {
    background: #026b5b;
}

/* =================================== */
/* 4. LOADING & EMPTY STATE */
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
/* 5. NEWS TABLE */
/* =================================== */
.news-table-container {
    background: #fff;
    border-radius: 16px;
    overflow: hidden;
    box-shadow: 0 2px 12px rgba(3, 114, 102, 0.07);
    border: 1px solid #eef0ef;
    overflow-x: auto;
}

.news-table {
    width: 100%;
    min-width: 800px;
    border-collapse: collapse;
}

.news-table thead {
    background: #f7f9f8;
}

.news-table th {
    padding: 16px;
    text-align: left;
    font-size: 14px;
    font-weight: 700;
    color: #555;
    text-transform: uppercase;
    letter-spacing: 0.5px;
    border-bottom: 2px solid #e5e5e5;
}

.news-table tbody tr {
    border-bottom: 1px solid #f0f0f0;
    transition: all 0.2s;
}

.news-table tbody tr:hover {
    background: #f9fafb;
}

.news-table tbody tr.selected {
    background: #f0f8f7;
}

.news-table td {
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
    min-width: 300px;
}

.title-cell {
    display: flex;
    gap: 14px;
    align-items: flex-start;
}

.news-thumbnail {
    width: 80px;
    height: 60px;
    border-radius: 8px;
    overflow: hidden;
    flex-shrink: 0;
    background: #f0f0f0;
}

.news-thumbnail img {
    width: 100%;
    height: 100%;
    object-fit: cover;
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
    overflow: hidden;
    text-overflow: ellipsis;
    white-space: nowrap;
}

.col-category {
    width: 160px;
}

.category-text {
    font-size: 14px;
    color: #555;
    font-weight: 500;
    white-space: nowrap;
}

.col-status {
    width: 140px;
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

.status-badge.published {
    background: #d1fae5;
    color: #065f46;
}

.status-badge.scheduled {
    background: #fef3c7;
    color: #92400e;
}

.status-badge.draft {
    background: #e5e7eb;
    color: #4b5563;
}

.status-dot {
    width: 8px;
    height: 8px;
    border-radius: 50%;
    background: currentColor;
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
    white-space: nowrap;
}

.date-cell svg {
    flex-shrink: 0;
    color: #999;
}

.col-actions {
    width: 140px;
    text-align: center;
}

.action-buttons {
    display: flex;
    gap: 6px;
    justify-content: center;
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
/* 6. BULK ACTIONS BAR */
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

.btn-bulk.btn-publish {
    background: #d1fae5;
    color: #065f46;
}

.btn-bulk.btn-publish:hover {
    background: #a7f3d0;
}

.btn-bulk.btn-delete-bulk {
    background: #fee2e2;
    color: #dc2626;
}

.btn-bulk.btn-delete-bulk:hover {
    background: #fecaca;
}

/* =================================== */
/* 7. PAGINATION */
/* =================================== */
.pagination {
    display: flex;
    justify-content: center;
    align-items: center;
    gap: 8px;
    margin-top: 30px;
    padding: 20px 0;
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
/* 8. SIDE PANEL ACTIONS */
/* =================================== */
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

/* =================================== */
/* 9. DELETE CONFIRMATION MODAL */
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

@keyframes fadeIn {
    from {
        opacity: 0;
    }

    to {
        opacity: 1;
    }
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
/* 10. DROPDOWN TRANSITIONS */
/* =================================== */
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

/* =================================== */
/* 11. RESPONSIVE */
/* =================================== */
@media (max-width: 1024px) {
    .news-table {
        font-size: 14px;
    }

    .col-title {
        min-width: 250px;
    }

    .news-thumbnail {
        width: 60px;
        height: 45px;
    }
}

@media (max-width: 768px) {
    .manage-news-page {
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

    .status-date-filters {
        flex-direction: column;
        align-items: stretch;
    }

    .date-filter-dropdown {
        width: 100%;
    }

    .bulk-actions-bar {
        flex-direction: column;
        align-items: stretch;
        max-width: calc(100% - 30px);
    }

    .bulk-buttons {
        flex-direction: column;
        width: 100%;
    }

    .btn-bulk {
        width: 100%;
        justify-content: center;
    }

    .modal-content {
        max-width: 100%;
        max-height: 100vh;
        border-radius: 0;
    }

    .delete-confirm-modal {
        padding: 20px;
    }

    .modal-actions {
        flex-direction: column;
    }

    .btn-cancel,
    .btn-confirm-delete {
        width: 100%;
    }
}

@media (max-width: 480px) {
    .stats-overview {
        grid-template-columns: 1fr;
    }

    .header-content h1 {
        font-size: 24px;
    }

    .date-inputs {
        flex-direction: column;
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