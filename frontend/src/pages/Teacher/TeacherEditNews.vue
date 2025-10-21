<!-- frontend/src/pages/Teacher/TeacherEditNews.vue -->
<template>
    <DashboardLayout role="teacher">
        <div class="edit-news-page">
            <div class="page-header">
                <button class="btn-back" @click="goBack">
                    <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none"
                        stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                        <line x1="19" y1="12" x2="5" y2="12"></line>
                        <polyline points="12 19 5 12 12 5"></polyline>
                    </svg>
                    ย้อนกลับ
                </button>
                <h1>แก้ไขข่าวสารและกิจกรรม</h1>
                <p class="page-subtitle">แก้ไขรายละเอียดข่าวประชาสัมพันธ์และกิจกรรม</p>
            </div>

            <div v-if="isLoading" class="loading-state">
                <div class="spinner"></div>
                <p>กำลังโหลดข้อมูล...</p>
            </div>

            <div v-else-if="!newsForm.id" class="error-state">
                <svg xmlns="http://www.w3.org/2000/svg" width="64" height="64" viewBox="0 0 24 24" fill="none"
                    stroke="currentColor" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round">
                    <circle cx="12" cy="12" r="10"></circle>
                    <line x1="15" y1="9" x2="9" y2="15"></line>
                    <line x1="9" y1="9" x2="15" y2="15"></line>
                </svg>
                <h3>ไม่พบข่าวสาร</h3>
                <p>ไม่สามารถโหลดข้อมูลข่าวสารได้</p>
                <button class="btn-back-error" @click="router.push({ name: 'TeacherManageNews' })">
                    กลับไปหน้าจัดการข่าวสาร
                </button>
            </div>

            <div v-else class="edit-content">
                <div class="section-card">
                    <h3>รูปภาพประกอบ</h3>
                    <p class="helper-text">อัปโหลดรูปภาพประกอบเพื่อดึงดูดความสนใจจากผู้อ่าน (ไม่เกิน 5 รูป)</p>

                    <div class="upload-area">
                        <div class="uploaded-images" v-if="uploadedImages.length > 0">
                            <div v-for="(image, index) in uploadedImages" :key="index" class="image-preview-item"
                                :class="{ active: currentImageIndex === index }" @click="currentImageIndex = index"
                                style="cursor: pointer;">
                                <img :src="image.preview" :alt="'Preview ' + (index + 1)" />
                                <button class="btn-remove-image" @click.stop="removeImage(index)">
                                    <svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24"
                                        fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round"
                                        stroke-linejoin="round">
                                        <line x1="18" y1="6" x2="6" y2="18"></line>
                                        <line x1="6" y1="6" x2="18" y2="18"></line>
                                    </svg>
                                </button>
                            </div>

                            <div v-if="uploadedImages.length < 5" class="upload-placeholder-small"
                                @click="triggerFileUpload">
                                <svg xmlns="http://www.w3.org/2000/svg" width="30" height="30" viewBox="0 0 24 24"
                                    fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round"
                                    stroke-linejoin="round">
                                    <path d="M21 15v4a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-4"></path>
                                    <polyline points="17 8 12 3 7 8"></polyline>
                                    <line x1="12" y1="3" x2="12" y2="15"></line>
                                </svg>
                                <span class="upload-small-text">เพิ่มรูป</span>
                            </div>
                        </div>

                        <div v-if="uploadedImages.length === 0" class="upload-placeholder" @click="triggerFileUpload">
                            <svg xmlns="http://www.w3.org/2000/svg" width="40" height="40" viewBox="0 0 24 24"
                                fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round"
                                stroke-linejoin="round">
                                <path d="M21 15v4a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-4"></path>
                                <polyline points="17 8 12 3 7 8"></polyline>
                                <line x1="12" y1="3" x2="12" y2="15"></line>
                            </svg>
                            <p>คลิกเพื่อเพิ่มรูปภาพประกอบ</p>
                            <span class="upload-hint">PNG, JPG สูงสุด 5MB</span>
                        </div>

                        <input ref="fileInput" type="file" accept="image/*" multiple @change="handleFileUpload"
                            hidden />
                    </div>

                    <div v-if="uploadedImages.length > 0" class="carousel-preview">
                        <h4>ภาพตัวอย่าง</h4>
                        <div class="carousel-container">
                            <img :src="uploadedImages[currentImageIndex].preview" alt="Preview"
                                class="carousel-image" />
                            <div class="carousel-controls">
                                <button class="btn-carousel" @click="prevImage" :disabled="uploadedImages.length <= 1">
                                    <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24"
                                        fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round"
                                        stroke-linejoin="round">
                                        <polyline points="15 18 9 12 15 6"></polyline>
                                    </svg>
                                </button>
                                <div class="carousel-dots">
                                    <span v-for="(img, idx) in uploadedImages" :key="idx" class="dot"
                                        :class="{ active: idx === currentImageIndex }"></span>
                                </div>
                                <button class="btn-carousel" @click="nextImage" :disabled="uploadedImages.length <= 1">
                                    <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24"
                                        fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round"
                                        stroke-linejoin="round">
                                        <polyline points="9 18 15 12 9 6"></polyline>
                                    </svg>
                                </button>
                            </div>
                        </div>
                    </div>
                </div>

                <div class="section-card">
                    <h3>เขียนข่าวสาร <span class="required">*</span></h3>
                    <p class="helper-text">แก้ไขรายละเอียดเกี่ยวกับข่าวสารหรือกิจกรรม</p>

                    <div class="form-group">
                        <label>หัวข้อข่าว</label>
                        <input type="text" v-model="newsForm.title"
                            placeholder="เช่น เปิดรับสมัครสหกิจศึกษา ภาคเรียนที่ 1/2568" class="form-input" />
                    </div>

                    <div class="form-group">
                        <label>หมวดหมู่</label>
                        <div class="custom-dropdown" ref="categoryDropdownRef"
                            @click="isCategoryDropdownOpen = !isCategoryDropdownOpen">
                            <div class="dropdown-selected">
                                <span>{{ newsForm.category || 'เลือกหมวดหมู่' }}</span>
                                <svg class="dropdown-arrow" :class="{ open: isCategoryDropdownOpen }"
                                    xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24"
                                    fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round"
                                    stroke-linejoin="round">
                                    <polyline points="6 9 12 15 18 9"></polyline>
                                </svg>
                            </div>
                            <transition name="dropdown-fade">
                                <ul v-if="isCategoryDropdownOpen" class="dropdown-menu">
                                    <li v-for="cat in categories" :key="cat" @click.stop="selectCategory(cat)">
                                        {{ cat }}
                                    </li>
                                </ul>
                            </transition>
                        </div>
                        <p class="field-hint">เลือกหมวดหมู่ที่เหมาะสมกับข่าวสารที่ต้องการประกาศ</p>
                    </div>

                    <div class="form-group">
                        <label>รายละเอียด</label>
                        <div class="editor-toolbar">
                            <button type="button" class="toolbar-btn" @click="execCommand('bold')" title="Bold">
                                <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24"
                                    fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round"
                                    stroke-linejoin="round">
                                    <path d="M6 4h8a4 4 0 0 1 4 4 4 4 0 0 1-4 4H6z"></path>
                                    <path d="M6 12h9a4 4 0 0 1 4 4 4 4 0 0 1-4 4H6z"></path>
                                </svg>
                            </button>
                            <button type="button" class="toolbar-btn" @click="execCommand('italic')" title="Italic">
                                <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24"
                                    fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round"
                                    stroke-linejoin="round">
                                    <line x1="19" y1="4" x2="10" y2="4"></line>
                                    <line x1="14" y1="20" x2="5" y2="20"></line>
                                    <line x1="15" y1="4" x2="9" y2="20"></line>
                                </svg>
                            </button>
                            <button type="button" class="toolbar-btn" @click="execCommand('underline')"
                                title="Underline">
                                <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24"
                                    fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round"
                                    stroke-linejoin="round">
                                    <path d="M6 3v7a6 6 0 0 0 6 6 6 6 0 0 0 6-6V3"></path>
                                    <line x1="4" y1="21" x2="20" y2="21"></line>
                                </svg>
                            </button>
                            <div class="toolbar-divider"></div>
                            <button type="button" class="toolbar-btn" @click="execCommand('insertUnorderedList')"
                                title="Bullet List">
                                <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24"
                                    fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round"
                                    stroke-linejoin="round">
                                    <line x1="8" y1="6" x2="21" y2="6"></line>
                                    <line x1="8" y1="12" x2="21" y2="12"></line>
                                    <line x1="8" y1="18" x2="21" y2="18"></line>
                                    <line x1="3" y1="6" x2="3.01" y2="6"></line>
                                    <line x1="3" y1="12" x2="3.01" y2="12"></line>
                                    <line x1="3" y1="18" x2="3.01" y2="18"></line>
                                </svg>
                            </button>
                            <button type="button" class="toolbar-btn" @click="execCommand('insertOrderedList')"
                                title="Numbered List">
                                <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24"
                                    fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round"
                                    stroke-linejoin="round">
                                    <line x1="10" y1="6" x2="21" y2="6"></line>
                                    <line x1="10" y1="12" x2="21" y2="12"></line>
                                    <line x1="10" y1="18" x2="21" y2="18"></line>
                                    <path d="M4 6h1v4"></path>
                                    <path d="M4 10h2"></path>
                                    <path d="M6 18H4c0-1 2-2 2-3s-1-1.5-2-1"></path>
                                </svg>
                            </button>
                        </div>
                        <div ref="editor" contenteditable="true" @input="updateDescription" class="form-editor"
                            placeholder="เขียนรายละเอียดข่าวสาร กิจกรรม และข้อมูลที่ต้องการประชาสัมพันธ์..."></div>
                    </div>
                </div>

                <div class="section-card">
                    <h3>กำหนดการประกาศ <span class="required">*</span></h3>
                    <p class="helper-text">แก้ไขสถานะและเวลาที่ต้องการประกาศข่าวสาร</p>

                    <div class="form-group">
                        <!-- แสดง Notice ถ้าข่าวสารเผยแพร่แล้ว -->
                        <div v-if="isPublished" class="locked-status-notice">
                            <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24"
                                fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round"
                                stroke-linejoin="round">
                                <rect x="3" y="11" width="18" height="11" rx="2" ry="2"></rect>
                                <path d="M7 11V7a5 5 0 0 1 10 0v4"></path>
                            </svg>
                            <p>ข่าวสารนี้ถูกเผยแพร่แล้ว ไม่สามารถเปลี่ยนสถานะการประกาศได้</p>
                        </div>

                        <!-- ตัวเลือกการประกาศ (เหมือน TeacherCreateNews) -->
                        <div class="publish-options">
                            <!-- ตัวเลือก 1: แบบร่าง -->
                            <label class="radio-item" :class="{ disabled: isPublished, active: publishType === 'draft' }">
                                <input type="radio" v-model="publishType" value="draft" name="publishType"
                                    :disabled="isPublished" />
                                <div class="radio-content">
                                    <div class="radio-box"></div>
                                    <div class="radio-info">
                                        <span class="radio-label">บันทึกเป็นแบบร่าง</span>
                                        <span class="radio-description">เก็บข่าวสารไว้โดยยังไม่เผยแพร่</span>
                                    </div>
                                </div>
                            </label>

                            <!-- ตัวเลือก 2: ประกาศทันที -->
                            <label class="radio-item" :class="{ disabled: isPublished, active: publishType === 'immediate' }">
                                <input type="radio" v-model="publishType" value="immediate" name="publishType"
                                    :disabled="isPublished" />
                                <div class="radio-content">
                                    <div class="radio-box"></div>
                                    <div class="radio-info">
                                        <span class="radio-label">ประกาศทันที</span>
                                        <span class="radio-description">ข่าวสารจะถูกเผยแพร่ทันทีหลังจากบันทึก</span>
                                    </div>
                                </div>
                            </label>

                            <!-- ตัวเลือก 3: กำหนดวันเวลา -->
                            <label class="radio-item" :class="{ disabled: isPublished, active: publishType === 'scheduled' }">
                                <input type="radio" v-model="publishType" value="scheduled" name="publishType"
                                    :disabled="isPublished" />
                                <div class="radio-content">
                                    <div class="radio-box"></div>
                                    <div class="radio-info">
                                        <span class="radio-label">กำหนดวันเวลาประกาศ</span>
                                        <span class="radio-description">เลือกวันและเวลาที่ต้องการให้ประกาศ</span>
                                    </div>
                                </div>
                            </label>
                        </div>

                        <!-- แสดงช่องกรอกวันที่เมื่อเลือก scheduled หรือ published -->
                        <div v-if="publishType === 'scheduled' || isPublished" class="schedule-inputs">
                            <div class="datetime-group">
                                <div class="input-wrapper">
                                    <label>วันที่ประกาศ</label>
                                    <input type="date" v-model="newsForm.publishDate" class="form-input" :min="minDate"
                                        :disabled="isPublished" />
                                </div>
                                <div class="input-wrapper">
                                    <label>เวลาที่ประกาศ</label>
                                    <input type="time" v-model="newsForm.publishTime" class="form-input"
                                        :disabled="isPublished" />
                                </div>
                            </div>
                        </div>
                    </div>
                </div>

                <div class="action-buttons">
                    <button class="btn-secondary" @click="handleCancel">ยกเลิก</button>
                    <button class="btn-primary" @click="handleSave" :disabled="!isFormValid || isSaving">
                        <svg v-if="!isSaving" xmlns="http://www.w3.org/2000/svg" width="18" height="18"
                            viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"
                            stroke-linecap="round" stroke-linejoin="round">
                            <path d="M19 21H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h11l5 5v11a2 2 0 0 1-2 2z"></path>
                            <polyline points="17 21 17 13 7 13 7 21"></polyline>
                            <polyline points="7 3 7 8 15 8"></polyline>
                        </svg>
                        <span>{{ isSaving ? 'กำลังบันทึก...' : 'บันทึกการแก้ไข' }}</span>
                    </button>
                </div>
            </div>
        </div>
    </DashboardLayout>
</template>

<script setup>
import { ref, reactive, computed, watch, onMounted, onBeforeUnmount } from 'vue';
import { useRouter, useRoute } from 'vue-router';
import DashboardLayout from '../../components/DashboardLayout.vue';
import { API_CONFIG } from '../../../config/api';

// =======================================================
// 1. ROUTER & ROUTE
// =======================================================
const router = useRouter();
const route = useRoute();

// =======================================================
// 2. CONSTANTS
// =======================================================
const categories = [
    "ประชาสัมพันธ์",
    "ข่าวด่วน",
    "กิจกรรม",
    "ประกาศผลการคัดเลือก"
];

// =======================================================
// 3. STATE MANAGEMENT
// =======================================================
const isLoading = ref(true);
const isSaving = ref(false);
const publishType = ref('immediate');
const originalStatus = ref(null);

const newsForm = reactive({
    id: null,
    title: '',
    category: '',
    description: '',
    status: 'published',
    publishDate: '',
    publishTime: ''
});

const uploadedImages = ref([]);
const currentImageIndex = ref(0);

const isCategoryDropdownOpen = ref(false);

// =======================================================
// 4. DOM REFS
// =======================================================
const fileInput = ref(null);
const editor = ref(null);
const categoryDropdownRef = ref(null);

// =======================================================
// 5. COMPUTED PROPERTIES
// =======================================================
const isPublished = computed(() => {
    return originalStatus.value === 'immediate' || originalStatus.value === 'published';
});

const minDate = computed(() => {
    const today = new Date();
    return today.toISOString().split('T')[0];
});

const isFormValid = computed(() => {
    const hasBasicInfo = newsForm.title.trim() !== '' &&
        newsForm.category.trim() !== '' &&
        newsForm.description.trim() !== '';
    
    // ถ้าเป็นแบบร่าง: ไม่บังคับวันที่
    if (publishType.value === 'draft') {
        return hasBasicInfo;
    }
    
    // ถ้าเป็นกำหนดวันเวลา: บังคับวันที่
    if (publishType.value === 'scheduled') {
        return hasBasicInfo && newsForm.publishDate !== '';
    }
    
    // ประกาศทันทีหรือเผยแพร่แล้ว
    return hasBasicInfo;
});

// =======================================================
// 6. LIFECYCLE - LOAD DATA
// =======================================================
onMounted(() => {
    document.addEventListener("click", handleClickOutside);
    loadNewsData();
});

onBeforeUnmount(() => {
    document.removeEventListener("click", handleClickOutside);
});

const loadNewsData = async () => {
    const newsId = parseInt(route.params.id);
    isLoading.value = true;

    try {
        const response = await fetch(`${API_CONFIG.baseURL}${API_CONFIG.endpoints.announcements}/${newsId}`);
        
        if (!response.ok) {
            throw new Error('Failed to fetch announcement');
        }

        const foundNews = await response.json();

        if (foundNews) {
            // Load news data
            newsForm.id = foundNews.post_id;
            newsForm.title = foundNews.title;
            newsForm.category = foundNews.category || 'ประชาสัมพันธ์';
            newsForm.description = foundNews.description;
            newsForm.status = foundNews.status || 'immediate';
            originalStatus.value = foundNews.status || 'immediate';

            // ตั้งค่า publishType ตามสถานะ
            if (foundNews.status === 'draft') {
                publishType.value = 'draft';
            } else if (foundNews.status === 'scheduled') {
                publishType.value = 'scheduled';
            } else {
                publishType.value = 'immediate';
            }

            // Handle date
            if (foundNews.publish_date) {
                const publishDateTime = new Date(foundNews.publish_date);
                newsForm.publishDate = publishDateTime.toISOString().split('T')[0];
                newsForm.publishTime = publishDateTime.toTimeString().slice(0, 5);
            } else if (foundNews.created_at) {
                const createdDate = new Date(foundNews.created_at);
                newsForm.publishDate = createdDate.toISOString().split('T')[0];
                newsForm.publishTime = '09:00';
            }

            // Load images
            if (foundNews.attachments && foundNews.attachments.length > 0) {
                uploadedImages.value = foundNews.attachments.map(attachment => ({
                    preview: `${API_CONFIG.baseURL}/uploads/news/${attachment.filename}`,
                    isExisting: true,
                    fileId: attachment.file_id
                }));
            }

            // Load description into editor
            if (editor.value) {
                editor.value.innerHTML = foundNews.description;
            }
        }
    } catch (error) {
        console.error('Error loading news data:', error);
        // แสดง error state
        newsForm.id = null;
    } finally {
        isLoading.value = false;
    }
};
// =======================================================
// 7. IMAGE HANDLING
// =======================================================
const triggerFileUpload = () => {
    fileInput.value.click();
};

const handleFileUpload = (event) => {
    const files = Array.from(event.target.files);
    const remainingSlots = 5 - uploadedImages.value.length;

    files.slice(0, remainingSlots).forEach(file => {
        if (file.size > 5 * 1024 * 1024) {
            alert(`ไฟล์ "${file.name}" มีขนาดใหญ่เกิน 5MB`);
            return;
        }
        const reader = new FileReader();
        reader.onload = (e) => {
            uploadedImages.value.push({
                file: file,
                preview: e.target.result,
                isExisting: false
            });
        };
        reader.readAsDataURL(file);
    });
    event.target.value = '';
};

const removeImage = (index) => {
    uploadedImages.value.splice(index, 1);
    if (currentImageIndex.value >= uploadedImages.value.length) {
        currentImageIndex.value = Math.max(0, uploadedImages.value.length - 1);
    }
};

const prevImage = () => {
    if (currentImageIndex.value > 0) {
        currentImageIndex.value--;
    }
};

const nextImage = () => {
    if (currentImageIndex.value < uploadedImages.value.length - 1) {
        currentImageIndex.value++;
    }
};

// =======================================================
// 8. DROPDOWN HANDLERS
// =======================================================
const selectCategory = (cat) => {
    newsForm.category = cat;
    isCategoryDropdownOpen.value = false;
};

const handleClickOutside = (event) => {
    if (categoryDropdownRef.value && !categoryDropdownRef.value.contains(event.target)) {
        isCategoryDropdownOpen.value = false;
    }
};

// =======================================================
// 9. RICH TEXT EDITOR
// =======================================================
const execCommand = (command) => {
    document.execCommand(command, false, null);
    editor.value.focus();
    updateDescription();
};

const updateDescription = () => {
    newsForm.description = editor.value.innerHTML;
};

// Watch publishType เพื่ออัปเดต newsForm.status (ถ้ายังไม่ published)
watch(publishType, (newType) => {
    // ถ้าข่าวสารยังไม่ published ให้เปลี่ยนสถานะได้
    if (!isPublished.value) {
        if (newType === 'draft') {
            newsForm.status = 'draft';
            // ล้างวันที่
            newsForm.publishDate = '';
            newsForm.publishTime = '';
        } else if (newType === 'scheduled') {
            newsForm.status = 'scheduled';
        } else if (newType === 'immediate') {
            newsForm.status = 'published';
            // ตั้งวันที่เป็นวันปัจจุบัน
            const today = new Date();
            newsForm.publishDate = today.toISOString().split('T')[0];
            newsForm.publishTime = '09:00';
        }
    }
});

// =======================================================
// 10. FORM ACTIONS
// =======================================================
const goBack = () => {
    router.back();
};

const handleCancel = () => {
    if (confirm('คุณต้องการยกเลิกการแก้ไขหรือไม่? การเปลี่ยนแปลงที่ยังไม่บันทึกจะหายไป')) {
        router.back();
    }
};

const handleSave = async () => {
    if (!isFormValid.value) {
        alert('กรุณากรอกข้อมูลให้ครบถ้วน');
        return;
    }

    isSaving.value = true;

    try {
        const formData = new FormData();
        formData.append('title', newsForm.title);
        formData.append('description', newsForm.description);
        formData.append('category', newsForm.category);
        
        // Map status
        const statusMap = {
            'draft': 'draft',
            'immediate': 'immediate',
            'scheduled': 'scheduled'
        };
        formData.append('status', statusMap[publishType.value]);

        // เพิ่ม publish_date ถ้าเป็น scheduled
        if (publishType.value === 'scheduled' && newsForm.publishDate) {
            formData.append('publish_date', `${newsForm.publishDate}T${newsForm.publishTime || '00:00'}`);
        }

        // เพิ่มรูปภาพใหม่
        uploadedImages.value.forEach(img => {
            if (img.file) {
                formData.append('files', img.file);
            }
        });

        const response = await fetch(
            `${API_CONFIG.baseURL}${API_CONFIG.endpoints.announcements}/${newsForm.id}`, 
            {
                method: 'PUT',
                body: formData
            }
        );

        if (!response.ok) {
            throw new Error('Failed to update announcement');
        }

        alert('บันทึกการแก้ไขเรียบร้อยแล้ว');
        router.push({ name: 'TeacherManageNews' });

    } catch (error) {
        console.error('Error updating announcement:', error);
        alert('เกิดข้อผิดพลาดในการบันทึก กรุณาลองใหม่อีกครั้ง');
    } finally {
        isSaving.value = false;
    }
};
</script>

<style scoped>
@import url("https://fonts.googleapis.com/css2?family=Kanit:wght@400;500;600;700&display=swap");

* {
    font-family: "Kanit", sans-serif;
    box-sizing: border-box;
}

.edit-news-page {
    max-width: 900px;
    margin: 0 auto;
    padding: 20px;
}

.required {
    color: #ef4444;
    font-size: 18px;
}

/* =================================== */
/* PAGE HEADER & NAVIGATION */
/* =================================== */
.page-header {
    margin-bottom: 30px;
}

.btn-back {
    display: inline-flex;
    align-items: center;
    gap: 8px;
    padding: 10px 20px;
    background: #f0f0f0;
    border: none;
    border-radius: 10px;
    font-size: 15px;
    font-weight: 600;
    color: #666;
    cursor: pointer;
    transition: all 0.2s;
    margin-bottom: 20px;
}

.btn-back:hover {
    background: #e0e0e0;
    color: #333;
}

.page-header h1 {
    font-size: 28px;
    font-weight: 700;
    color: #222;
    margin: 0 0 10px 0;
}

.page-subtitle {
    font-size: 15px;
    color: #666;
    margin: 0;
    line-height: 1.6;
}

/* =================================== */
/* LOADING & ERROR STATES */
/* =================================== */
.loading-state,
.error-state {
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

.error-state svg {
    color: #dc2626;
}

.error-state h3 {
    font-size: 20px;
    font-weight: 600;
    color: #555;
    margin: 0;
}

.error-state p {
    font-size: 15px;
    color: #888;
    margin: 0;
}

.btn-back-error {
    margin-top: 12px;
    padding: 12px 24px;
    background: #037266;
    color: #fff;
    border: none;
    border-radius: 12px;
    font-size: 15px;
    font-weight: 600;
    cursor: pointer;
    transition: all 0.2s;
}

.btn-back-error:hover {
    background: #026b5b;
    transform: translateY(-2px);
}

/* =================================== */
/* CARD & HELPERS */
/* =================================== */
.edit-content {
    animation: fadeIn 0.3s ease;
}

@keyframes fadeIn {
    from {
        opacity: 0;
        transform: translateY(10px);
    }

    to {
        opacity: 1;
        transform: translateY(0);
    }
}

.section-card {
    background: #fff;
    border-radius: 16px;
    padding: 30px;
    box-shadow: 0 2px 12px rgba(3, 114, 102, 0.07);
    border: 1px solid #eef0ef;
    margin-bottom: 25px;
}

.section-card h3 {
    font-size: 20px;
    font-weight: 700;
    color: #222;
    margin: 0 0 10px 0;
    display: flex;
    align-items: center;
    gap: 6px;
}

.helper-text {
    font-size: 14px;
    color: #777;
    margin-bottom: 25px;
    line-height: 1.6;
}

.field-hint {
    font-size: 13px;
    color: #999;
    margin-top: 8px;
    line-height: 1.5;
}

/* =================================== */
/* IMAGE UPLOAD AREA */
/* =================================== */
.upload-area {
    margin-bottom: 25px;
}

.uploaded-images {
    display: grid;
    grid-template-columns: repeat(auto-fill, minmax(140px, 1fr));
    gap: 15px;
    margin-bottom: 15px;
}

.image-preview-item {
    position: relative;
    aspect-ratio: 1;
    border-radius: 12px;
    overflow: hidden;
    border: 2px solid #e5e5e5;
    transition: all 0.2s;
}

.image-preview-item:hover {
    border-color: #037266;
    transform: scale(1.02);
}

.image-preview-item.active {
    border-color: #037266;
    box-shadow: 0 0 0 3px rgba(3, 114, 102, 0.2);
    transform: scale(1.05);
}

.image-preview-item img {
    width: 100%;
    height: 100%;
    object-fit: cover;
}

.btn-remove-image {
    position: absolute;
    top: 8px;
    right: 8px;
    width: 32px;
    height: 32px;
    border-radius: 50%;
    background: rgba(239, 68, 68, 0.9);
    border: none;
    color: #fff;
    cursor: pointer;
    display: flex;
    align-items: center;
    justify-content: center;
    transition: all 0.2s;
    backdrop-filter: blur(4px);
}

.btn-remove-image:hover {
    background: #dc2626;
    transform: scale(1.1);
}

.upload-placeholder {
    border: 2px dashed #ccc;
    border-radius: 12px;
    padding: 40px 30px;
    text-align: center;
    background: #fafafa;
    cursor: pointer;
    transition: all 0.2s;
}

.upload-placeholder:hover {
    border-color: #037266;
    background: #f0f8f7;
}

.upload-placeholder svg {
    color: #999;
    margin-bottom: 15px;
}

.upload-placeholder p {
    font-size: 15px;
    color: #666;
    margin: 0 0 8px 0;
    font-weight: 500;
}

.upload-hint {
    font-size: 13px;
    color: #999;
}

.upload-placeholder-small {
    aspect-ratio: 1;
    border: 2px dashed #ccc;
    border-radius: 12px;
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    cursor: pointer;
    transition: all 0.2s;
    background: #fafafa;
}

.upload-placeholder-small:hover {
    border-color: #037266;
    background: #f0f8f7;
}

.upload-small-text {
    font-size: 13px;
    color: #666;
    margin-top: 6px;
    font-weight: 500;
}

/* =================================== */
/* CAROUSEL PREVIEW */
/* =================================== */
.carousel-preview {
    margin-top: 30px;
    padding-top: 30px;
    border-top: 2px solid #f0f0f0;
}

.carousel-preview h4 {
    font-size: 16px;
    font-weight: 600;
    color: #333;
    margin: 0 0 15px 0;
}

.carousel-container {
    position: relative;
    background: #f9f9f9;
    border-radius: 12px;
    padding: 20px;
}

.carousel-image {
    width: 100%;
    height: 400px;
    object-fit: contain;
    border-radius: 10px;
    background: #fff;
}

.carousel-controls {
    display: flex;
    align-items: center;
    justify-content: center;
    gap: 20px;
    margin-top: 20px;
}

.btn-carousel {
    width: 40px;
    height: 40px;
    border-radius: 50%;
    background: #037266;
    border: none;
    color: #fff;
    cursor: pointer;
    display: flex;
    align-items: center;
    justify-content: center;
    transition: all 0.2s;
}

.btn-carousel:hover:not(:disabled) {
    background: #026b5b;
    transform: scale(1.05);
}

.btn-carousel:disabled {
    background: #ccc;
    cursor: not-allowed;
    opacity: 0.5;
}

.carousel-dots {
    display: flex;
    gap: 8px;
}

.dot {
    width: 10px;
    height: 10px;
    border-radius: 50%;
    background: #d0d0d0;
    transition: all 0.2s;
}

.dot.active {
    background: #037266;
    width: 24px;
    border-radius: 5px;
}

/* =================================== */
/* FORM ELEMENTS */
/* =================================== */
.form-group {
    margin-bottom: 25px;
}

.form-group label {
    display: block;
    font-size: 15px;
    font-weight: 600;
    color: #333;
    margin-bottom: 10px;
}

.form-input {
    width: 100%;
    padding: 14px 16px;
    border: 2px solid #e5e5e5;
    border-radius: 10px;
    font-size: 15px;
    color: #333;
    transition: all 0.2s;
    font-family: "Kanit", sans-serif;
}

.form-input:focus {
    outline: none;
    border-color: #037266;
    background: #f0f8f7;
}

.form-input:disabled {
    background: #f5f5f5;
    color: #999;
    cursor: not-allowed;
}

/* =================================== */
/* CUSTOM DROPDOWN */
/* =================================== */
.custom-dropdown {
    position: relative;
    width: 100%;
    min-width: 250px;
    cursor: pointer;
    user-select: none;
}

.dropdown-selected {
    display: flex;
    align-items: center;
    justify-content: space-between;
    padding: 14px 16px;
    background: #fff;
    border: 2px solid #e5e5e5;
    border-radius: 10px;
    font-size: 15px;
    color: #333;
    font-weight: 500;
    transition: all 0.2s ease;
    box-shadow: none;
}

.dropdown-selected:hover {
    border-color: #037266;
    background: #f0f8f7;
    box-shadow: none;
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
    padding: 10px 12px;
    font-size: 15px;
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

/* =================================== */
/* RICH TEXT EDITOR */
/* =================================== */
.editor-toolbar {
    display: flex;
    align-items: center;
    gap: 4px;
    padding: 10px;
    background: #f9f9f9;
    border: 2px solid #e5e5e5;
    border-bottom: none;
    border-radius: 10px 10px 0 0;
}

.toolbar-btn {
    width: 36px;
    height: 36px;
    border: none;
    background: transparent;
    border-radius: 6px;
    display: flex;
    align-items: center;
    justify-content: center;
    cursor: pointer;
    color: #666;
    transition: all 0.2s;
}

.toolbar-btn:hover {
    background: #e0e0e0;
    color: #333;
}

.toolbar-divider {
    width: 1px;
    height: 24px;
    background: #d0d0d0;
    margin: 0 4px;
}

.form-editor {
    min-height: 300px;
    padding: 16px;
    border: 2px solid #e5e5e5;
    border-radius: 0 0 10px 10px;
    font-size: 15px;
    color: #333;
    line-height: 1.8;
    transition: all 0.2s;
    background: #fff;
    overflow-y: auto;
}

.form-editor:focus {
    outline: none;
    border-color: #037266;
    background: #f0f8f7;
}

.form-editor:empty:before {
    content: attr(placeholder);
    color: #999;
}

/* =================================== */
/* LOCKED STATUS NOTICE */
/* =================================== */
.locked-status-notice {
    background: #fff9e6;
    border: 1px solid #ffd43b;
    border-radius: 12px;
    padding: 16px 18px;
    display: flex;
    align-items: center;
    gap: 12px;
    margin-bottom: 20px;
}

.locked-status-notice svg {
    color: #f59e0b;
    flex-shrink: 0;
}

.locked-status-notice p {
    font-size: 14px;
    color: #92400e;
    margin: 0;
    font-weight: 500;
    line-height: 1.5;
}

/* =================================== */
/* PUBLISH OPTIONS (RADIO) */
/* =================================== */
.publish-options {
    display: flex;
    flex-direction: column;
    gap: 15px;
    margin-bottom: 20px;
}

.radio-item {
    display: block;
    cursor: pointer;
}

.radio-item.disabled {
    opacity: 0.5;
    cursor: not-allowed;
}

.radio-item input[type="radio"] {
    display: none;
}

.radio-content {
    display: flex;
    align-items: flex-start;
    gap: 12px;
    padding: 16px;
    background: #fff;
    border: 2px solid #e5e5e5;
    border-radius: 10px;
    transition: all 0.2s;
}

.radio-item:not(.disabled):hover .radio-content {
    border-color: #037266;
    background: #f0f8f7;
}

.radio-item input:checked + .radio-content {
    border-color: #037266;
    background: #f0f8f7;
}

.radio-item.disabled .radio-content {
    background: #f5f5f5;
    border-color: #d0d0d0;
    cursor: not-allowed;
}

.radio-box {
    width: 22px;
    height: 22px;
    border: 2px solid #d0d0d0;
    border-radius: 50%;
    position: relative;
    flex-shrink: 0;
    margin-top: 2px;
    transition: all 0.2s;
}

.radio-item input:checked + .radio-content .radio-box {
    border-color: #037266;
    background: #fff;
}

.radio-item.disabled .radio-box {
    border-color: #d0d0d0;
    background: #f5f5f5;
}

.radio-item input:checked + .radio-content .radio-box::after {
    content: '';
    position: absolute;
    top: 50%;
    left: 50%;
    transform: translate(-50%, -50%);
    width: 10px;
    height: 10px;
    border-radius: 50%;
    background: #037266;
}

.radio-item.disabled input:checked + .radio-content .radio-box::after {
    background: #999;
}

.radio-info {
    display: flex;
    flex-direction: column;
    gap: 4px;
    flex: 1;
}

.radio-label {
    font-size: 15px;
    font-weight: 600;
    color: #333;
}

.radio-item.disabled .radio-label {
    color: #999;
}

.radio-description {
    font-size: 13px;
    color: #777;
    line-height: 1.5;
}

.radio-item.disabled .radio-description {
    color: #aaa;
}

/* =================================== */
/* SCHEDULE INPUTS */
/* =================================== */
.schedule-inputs {
    margin-top: 20px;
    padding: 20px;
    background: #f9f9f9;
    border-radius: 10px;
    border: 1px solid #e5e5e5;
}

.datetime-group {
    display: grid;
    grid-template-columns: 1fr 1fr;
    gap: 15px;
}

.input-wrapper {
    display: flex;
    flex-direction: column;
    gap: 8px;
}

.input-wrapper label {
    font-size: 14px;
    font-weight: 600;
    color: #555;
    margin: 0;
}

/* =================================== */
/* ACTION BUTTONS */
/* =================================== */
.action-buttons {
    display: flex;
    justify-content: flex-end;
    gap: 15px;
    margin-top: 30px;
    padding-top: 25px;
    border-top: 1px solid #eee;
}

.btn-secondary,
.btn-primary {
    padding: 13px 32px;
    border: none;
    border-radius: 12px;
    font-size: 16px;
    font-weight: 600;
    cursor: pointer;
    transition: all 0.2s;
    display: flex;
    align-items: center;
    gap: 8px;
}

.btn-secondary {
    background: #f0f0f0;
    color: #666;
}

.btn-secondary:hover {
    background: #e0e0e0;
    color: #333;
}

.btn-primary {
    background: linear-gradient(135deg, #037266 0%, #026b5b 100%);
    color: #fff;
    box-shadow: 0 4px 12px rgba(3, 114, 102, 0.2);
}

.btn-primary:hover:not(:disabled) {
    background: linear-gradient(135deg, #026b5b 0%, #025a4d 100%);
    transform: translateY(-2px);
    box-shadow: 0 6px 16px rgba(3, 114, 102, 0.3);
}

.btn-primary:disabled {
    background: #ccc;
    cursor: not-allowed;
    box-shadow: none;
    opacity: 0.6;
}

/* =================================== */
/* RESPONSIVE */
/* =================================== */
@media (max-width: 768px) {
    .edit-news-page {
        padding: 15px;
    }

    .page-header h1 {
        font-size: 24px;
    }

    .section-card {
        padding: 20px;
    }

    .uploaded-images {
        grid-template-columns: repeat(auto-fill, minmax(100px, 1fr));
    }

    .carousel-image {
        height: 250px;
    }

    .action-buttons {
        flex-direction: column-reverse;
    }

    .btn-secondary,
    .btn-primary {
        width: 100%;
        justify-content: center;
    }

    .datetime-group {
        grid-template-columns: 1fr;
    }
}

@media (max-width: 480px) {
    .page-header h1 {
        font-size: 22px;
    }

    .uploaded-images {
        grid-template-columns: repeat(2, 1fr);
    }
}
</style>