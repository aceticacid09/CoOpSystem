<!-- frontend/src/pages/Company/CompanyCreateJobs.vue -->
<template>
    <DashboardLayout role="company">
        <div class="create-job-page">
            <!-- Header -->
            <div class="page-header">
                <button class="btn-back" @click="goBack">
                    <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none"
                        stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                        <line x1="19" y1="12" x2="5" y2="12"></line>
                        <polyline points="12 19 5 12 12 5"></polyline>
                    </svg>
                    ย้อนกลับ
                </button>
                <h1>สร้างประกาศงานใหม่</h1>
                <p class="page-subtitle">สร้างประกาศรับสมัครงานเพื่อค้นหาบุคลากรที่เหมาะสมกับองค์กรของคุณ</p>
            </div>

            <!-- Progress Steps -->
            <div class="progress-steps">
                <div class="step" :class="{ active: currentStep >= 1, completed: currentStep > 1 }">
                    <div class="step-number">1</div>
                    <div class="step-label">สร้างประกาศ</div>
                </div>
                <div class="step-line" :class="{ active: currentStep > 1 }"></div>
                <div class="step" :class="{ active: currentStep >= 2, completed: currentStep > 2 }">
                    <div class="step-number">2</div>
                    <div class="step-label">พรีวิว</div>
                </div>
                <div class="step-line" :class="{ active: currentStep > 2 }"></div>
                <div class="step" :class="{ active: currentStep >= 3 }">
                    <div class="step-number">3</div>
                    <div class="step-label">ตรวจสอบและประกาศ</div>
                </div>
            </div>

            <!-- Step 1: Create Job Posting -->
            <div v-if="currentStep === 1" class="step-content">
                <!-- Upload Images Section -->
                <div class="section-card">
                    <h3>รูปโปสเตอร์</h3>
                    <p class="helper-text">อัปโหลดรูปภาพประกอบเพื่อดึงดูดความสนใจจากผู้สมัครงาน (ไม่เกิน 5 รูป)</p>

                    <div class="upload-area">
                        <div class="uploaded-images" v-if="uploadedImages.length > 0">
                            <div v-for="(image, index) in uploadedImages" :key="index" class="image-preview-item">
                                <img :src="image.preview" alt="Preview" />
                                <button class="btn-remove-image" @click="removeImage(index)">
                                    <svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24"
                                        fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round"
                                        stroke-linejoin="round">
                                        <line x1="18" y1="6" x2="6" y2="18"></line>
                                        <line x1="6" y1="6" x2="18" y2="18"></line>
                                    </svg>
                                </button>
                            </div>

                            <!-- Placeholder สำหรับอัพโหลดเพิ่มเติม -->
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
                            <p>คลิกเพื่ออัปโหลดรูปภาพ</p>
                            <span class="upload-hint">PNG, JPG สูงสุด 5MB</span>
                        </div>

                        <input ref="fileInput" type="file" accept="image/*" multiple @change="handleFileUpload"
                            hidden />
                    </div>

                    <!-- Preview Carousel -->
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

                <!-- Job Details Section -->
                <div class="section-card">
                    <h3>เขียนประกาศงาน <span class="required">*</span></h3>
                    <p class="helper-text">เขียนรายละเอียดเกี่ยวกับตำแหน่งงานและคุณสมบัติผู้สมัคร</p>

                    <div class="form-group">
                        <label>ตำแหน่งงาน <span class="required">*</span></label>
                        <input type="text" v-model="jobForm.title" placeholder="เช่น Frontend Developer Intern"
                            class="form-input" />
                    </div>

                    <div class="form-group">
                        <label>ภาควิชา</label>
                        <div class="select-wrapper">
                            <select v-model="jobForm.department" class="form-select">
                                <option value="">เลือกภาควิชา/โปรแกรม</option>
                                <option value="computer">ภาควิชาคอมพิวเตอร์</option>
                                <option value="engineering">ภาควิชาวิศวกรรม</option>
                                <option value="science">ภาควิชาวิทยาศาสตร์</option>
                                <option value="business">ภาควิชาบริหารธุรกิจ</option>
                            </select>
                            <svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24"
                                fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round"
                                stroke-linejoin="round" class="select-icon">
                                <polyline points="6 9 12 15 18 9"></polyline>
                            </svg>
                        </div>
                        <p class="field-hint">เลือกภาควิชาที่ต้องการกำหนดให้ผู้สมัครงาน</p>
                    </div>

                    <div class="form-group">
                        <label>รายละเอียด <span class="required">*</span></label>
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
                            placeholder="เขียนรายละเอียดงาน คุณสมบัติผู้สมัคร และสิทธิประโยชน์..."></div>
                    </div>
                </div>

                <!-- Application Period Section -->
                <div class="section-card">
                    <h3>กำหนดระยะเวลารับสมัครงาน <span class="required">*</span></h3>
                    <p class="helper-text">กำหนดช่วงเวลาที่เปิดรับสมัครงานและระยะเวลาการพิจารณา</p>

                    <!-- HTML Template -->
<div class="form-group">
    <div class="period-header">
        <label>ระยะเวลารับสมัครงาน <span class="required">*</span></label>
        <div class="tooltip-wrapper">
            <button type="button" class="tooltip-btn" @mouseenter="showTooltip = true" @mouseleave="showTooltip = false">
                <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                    <circle cx="12" cy="12" r="10"></circle>
                    <line x1="12" y1="8" x2="12" y2="12"></line>
                    <line x1="12" y1="16" x2="12.01" y2="16"></line>
                </svg>
            </button>
            <div v-if="showTooltip" class="tooltip-content">
                <p><strong>คำแนะนำเกี่ยวกับระยะเวลารับสมัคร:</strong></p>
                <ul>
                    <li>7 วัน: เหมาะสำหรับตำแหน่งที่ต้องการเร่งด่วน</li>
                    <li>15 วัน: เหมาะสำหรับตำแหน่งงานทั่วไป</li>
                    <li>20 วัน: ระบบจะแสดงประกาศงานของคุณนานขึ้น เพิ่มโอกาสหาผู้สมัครที่เหมาะสม</li>
                </ul>
                <p class="notice-highlight">หมายเหตุ: ไม่มีจำกัดระยะเวลา กรอกเลขที่ต้องการได้</p>
            </div>
        </div>
    </div>

    <div class="period-dropdown-wrapper">
        <div class="period-select-container">
            <select v-model="selectedDateOption" class="period-select">
                <option value="7">7 วัน</option>
                <option value="15">15 วัน</option>
                <option value="20">20 วัน</option>
                <option value="custom">กำหนดเอง</option>
            </select>
            <svg class="select-chevron" xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                <polyline points="6 9 12 15 18 9"></polyline>
            </svg>
        </div>

        <!-- Custom Days Input -->
        <div v-if="selectedDateOption === 'custom'" class="custom-input-wrapper">
            <input type="number" v-model.number="customDays" placeholder="กรอกจำนวนวัน" class="custom-days-input" min="1" />
            <span class="days-suffix">วัน</span>
        </div>
    </div>
</div>
                </div>

                <!-- Email Notification Section -->
                <div class="section-card">
                    <h3>วิธีรับข้อมูลการสมัคร</h3>
                    <p class="helper-text">ระบุอีเมลที่ต้องการรับการแจ้งเตือนเมื่อมีผู้สมัครงานผ่านระบบ Science
                        Cooperative
                        Education System</p>

                    <div class="checkbox-group">
                        <label class="checkbox-item">
                            <input type="checkbox" v-model="jobForm.sendToCompanyEmail" />
                            <div class="checkbox-content">
                                <div class="checkbox-box"></div>
                                <span>ส่งการแจ้งเตือนไปยังอีเมลองค์กร</span>
                            </div>
                        </label>
                    </div>

                    <div v-if="jobForm.sendToCompanyEmail" class="form-group">
                        <input type="email" v-model="jobForm.notificationEmail" placeholder="example@gmail.com"
                            class="form-input" />
                        <p class="field-hint">ไม่บังคับหรือเว้นว่างเพื่อรับการแจ้งเตือน</p>
                    </div>
                </div>

                <!-- Action Buttons -->
                <div class="action-buttons">
                    <button class="btn-secondary" @click="cancelCreation">ยกเลิก</button>
                    <button class="btn-primary" @click="nextStep" :disabled="!isStep1Valid">ถัดไป</button>
                </div>
            </div>

            <!-- Step 2: Preview -->
            <div v-if="currentStep === 2" class="step-content">
                <div class="preview-card">
                    <h3>ตัวอย่างประกาศงาน</h3>
                    <p class="helper-text">ตรวจสอบรายละเอียดก่อนเผยแพร่ประกาศงาน</p>

                    <!-- Job Preview -->
                    <div class="job-preview">
                        <div class="preview-images" v-if="uploadedImages.length > 0">
                            <img :src="uploadedImages[0].preview" alt="Job Preview" class="preview-main-image" />
                            <div class="preview-image-count" v-if="uploadedImages.length > 1">
                                <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24"
                                    fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round"
                                    stroke-linejoin="round">
                                    <rect x="3" y="3" width="18" height="18" rx="2" ry="2"></rect>
                                    <circle cx="8.5" cy="8.5" r="1.5"></circle>
                                    <polyline points="21 15 16 10 5 21"></polyline>
                                </svg>
                                <span>+{{ uploadedImages.length - 1 }} รูป</span>
                            </div>
                        </div>

                        <div class="preview-content">
                            <h2 class="preview-title">{{ jobForm.title || 'ตำแหน่งงาน' }}</h2>
                            <p class="preview-company">บริษัท ABC จำกัด • {{ getDepartmentName(jobForm.department) }}
                            </p>

                            <div class="preview-meta">
                                <div class="meta-item">
                                    <svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24"
                                        fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round"
                                        stroke-linejoin="round">
                                        <rect x="3" y="4" width="18" height="18" rx="2" ry="2"></rect>
                                        <line x1="16" y1="2" x2="16" y2="6"></line>
                                        <line x1="8" y1="2" x2="8" y2="6"></line>
                                        <line x1="3" y1="10" x2="21" y2="10"></line>
                                    </svg>
                                    <span>{{ getDateRangePreview() }}</span>
                                </div>
                            </div>

                            <div class="preview-description">
                                <h4>รายละเอียดงาน</h4>
                                <div v-html="jobForm.description || 'ไม่มีรายละเอียด'"></div>
                            </div>
                        </div>
                    </div>
                </div>

                <div class="action-buttons">
                    <button class="btn-secondary" @click="prevStep">ย้อนกลับ</button>
                    <button class="btn-primary" @click="nextStep">ถัดไป</button>
                </div>
            </div>

            <!-- Step 3: Confirm and Publish -->
            <div v-if="currentStep === 3" class="step-content">
                <div class="confirmation-card">
                    <div class="confirmation-icon">
                        <svg xmlns="http://www.w3.org/2000/svg" width="48" height="48" viewBox="0 0 24 24" fill="none"
                            stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                            <path d="M22 11.08V12a10 10 0 1 1-5.93-9.14"></path>
                            <polyline points="22 4 12 14.01 9 11.01"></polyline>
                        </svg>
                    </div>
                    <h2>ตรวจสอบและยืนยันการประกาศ</h2>
                    <p class="confirmation-text">คุณต้องการเผยแพร่ประกาศงาน <strong>{{ jobForm.title }}</strong>
                        ใช่หรือไม่?</p>

                    <div class="confirmation-details">
                        <div class="detail-item">
                            <span class="detail-label">ตำแหน่ง:</span>
                            <span class="detail-value">{{ jobForm.title }}</span>
                        </div>
                        <div class="detail-item">
                            <span class="detail-label">ภาควิชา:</span>
                            <span class="detail-value">{{ getDepartmentName(jobForm.department) }}</span>
                        </div>
                        <div class="detail-item">
                            <span class="detail-label">ระยะเวลารับสมัคร:</span>
                            <span class="detail-value">{{ getDateRangePreview() }}</span>
                        </div>
                        <div class="detail-item">
                            <span class="detail-label">อีเมลรับการแจ้งเตือน:</span>
                            <span class="detail-value">{{ jobForm.notificationEmail || 'ใช้อีเมลองค์กร' }}</span>
                        </div>
                    </div>

                    <div class="confirmation-notice">
                        <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none"
                            stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                            <circle cx="12" cy="12" r="10"></circle>
                            <line x1="12" y1="8" x2="12" y2="12"></line>
                            <line x1="12" y1="16" x2="12.01" y2="16"></line>
                        </svg>
                        <p>เมื่อกดยืนยันแล้ว ประกาศงานของคุณจะถูกเผยแพร่ในระบบทันที และนักศึกษาสามารถดูและสมัครได้</p>
                    </div>
                </div>

                <div class="action-buttons">
                    <button class="btn-secondary" @click="prevStep">ย้อนกลับ</button>
                    <button class="btn-primary btn-confirm" @click="publishJob">ยืนยันและประกาศ</button>
                </div>
            </div>
        </div>
    </DashboardLayout>
</template>

<script setup>
import { ref, reactive, computed } from 'vue';
import { useRouter } from 'vue-router';
import DashboardLayout from '../../components/DashboardLayout.vue';

const router = useRouter();

// Current step
const currentStep = ref(1);

// File input ref
const fileInput = ref(null);

// Editor ref
const editor = ref(null);

// Uploaded images
const uploadedImages = ref([]);
const currentImageIndex = ref(0);

const showTooltip = ref(false);
const customDays = ref('');

// Date options
const dateOptions = [
    { value: '7', label: '7 วัน' },
    { value: '15', label: '15 วัน' },
    { value: '20', label: '20 วัน' },
    { value: 'custom', label: 'กำหนดเอง' }
];

const selectedDateOption = ref('15');

// Job form data
const jobForm = reactive({
    title: '',
    department: '',
    description: '',
    startDate: '',
    endDate: '',
    sendToCompanyEmail: false,
    notificationEmail: '',
    applicationDays: 15 // เพิ่มฟิลด์นี้
});

// Computed
const isStep1Valid = computed(() => {
    return jobForm.title.trim() !== '' && jobForm.description.trim() !== '';
});

// Methods
const goBack = () => {
    router.back();
};

const triggerFileUpload = () => {
    fileInput.value.click();
};

const handleFileUpload = (event) => {
    const files = Array.from(event.target.files);
    const remainingSlots = 5 - uploadedImages.value.length;

    files.slice(0, remainingSlots).forEach(file => {
        if (file.size > 5 * 1024 * 1024) {
            alert('ไฟล์มีขนาดใหญ่เกิน 5MB');
            return;
        }

        const reader = new FileReader();
        reader.onload = (e) => {
            uploadedImages.value.push({
                file: file,
                preview: e.target.result
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

const execCommand = (command) => {
    document.execCommand(command, false, null);
    editor.value.focus();
};

const updateDescription = () => {
    jobForm.description = editor.value.innerHTML;
};

const getDepartmentName = (dept) => {
    const departments = {
        computer: 'ภาควิชาคอมพิวเตอร์',
        engineering: 'ภาควิชาวิศวกรรม',
        science: 'ภาควิชาวิทยาศาสตร์',
        business: 'ภาควิชาบริหารธุรกิจ'
    };
    return departments[dept] || 'ไม่ระบุภาควิชา';
};

const getDateRangePreview = () => {
    if (selectedDateOption.value === 'custom') {
        if (jobForm.startDate && jobForm.endDate) {
            return `${formatDate(today)} - ${formatDate(endDate)} (${days} วัน)`;
        };
    } else {
        return dateOptions.find(option => option.value === selectedDateOption.value).label;
    }
};

const formatDate = (date) => {
    const d = typeof date === 'string' ? new Date(date) : date;
    const day = d.getDate();
    const months = [
        "ม.ค.", "ก.พ.", "มี.ค.", "เม.ย.", "พ.ค.", "มิ.ย.",
        "ก.ค.", "ส.ค.", "ก.ย.", "ต.ค.", "พ.ย.", "ธ.ค."
    ];
    const month = months[d.getMonth()];
    const year = d.getFullYear() + 543;
    return `${day} ${month} ${year}`;
};

const nextStep = () => {
    if (currentStep.value < 3) {
        currentStep.value++;
        window.scrollTo({ top: 0, behavior: 'smooth' });
    }
};

const prevStep = () => {
    if (currentStep.value > 1) {
        currentStep.value--;
        window.scrollTo({ top: 0, behavior: 'smooth' });
    }
};

const cancelCreation = () => {
    if (confirm('คุณต้องการยกเลิกการสร้างประกาศงานใช่หรือไม่?')) {
        router.back();
    }
};

const publishJob = () => {
    // Submit job to API
    console.log('Publishing job:', jobForm);
    alert('ประกาศงานสำเร็จ!');
    router.push({ name: 'CompanyJobs' }); // Adjust route name as needed
};
</script>

<style scoped>
@import url("https://fonts.googleapis.com/css2?family=Kanit:wght@400;500;600;700&display=swap");

* {
    font-family: "Kanit", sans-serif;
    box-sizing: border-box;
}

.create-job-page {
    max-width: 900px;
    margin: 0 auto;
    padding: 20px;
}

/* ========== PAGE HEADER ========== */
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

/* ========== PROGRESS STEPS ========== */
.progress-steps {
    display: flex;
    align-items: center;
    justify-content: space-between;
    margin-bottom: 40px;
    padding: 0 20px;
}

.step {
    display: flex;
    flex-direction: column;
    align-items: center;
    gap: 12px;
    flex: 0 0 auto;
}

.step-number {
    width: 50px;
    height: 50px;
    border-radius: 50%;
    background: #f0f0f0;
    color: #999;
    display: flex;
    align-items: center;
    justify-content: center;
    font-size: 20px;
    font-weight: 700;
    transition: all 0.3s;
    border: 3px solid #f0f0f0;
}

.step.active .step-number {
    background: #037266;
    color: #fff;
    border-color: #037266;
}

.step.completed .step-number {
    background: #10b981;
    color: #fff;
    border-color: #10b981;
}

.step-label {
    font-size: 14px;
    color: #999;
    font-weight: 500;
    text-align: center;
    transition: all 0.3s;
}

.step.active .step-label {
    color: #037266;
    font-weight: 600;
}

.step-line {
    flex: 1;
    height: 3px;
    background: #f0f0f0;
    transition: all 0.3s;
    margin: 0 20px;
    margin-bottom: 35px;
}

.step-line.active {
    background: #037266;
}

/* ========== SECTION CARD ========== */
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

.required {
    color: #ef4444;
    font-size: 18px;
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

/* ========== UPLOAD AREA ========== */
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
    background: rgba(220, 38, 38, 1);
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

.upload-placeholder-small svg {
    color: #999;
    margin-bottom: 8px;
}

.upload-small-text {
    font-size: 13px;
    color: #666;
    font-weight: 500;
}

/* ========== CAROUSEL PREVIEW ========== */
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
    transform: scale(1.1);
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

/* ========== FORM GROUPS ========== */
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

.form-input::placeholder {
    color: #999;
}

/* Chrome, Safari and Opera */
.form-input[type=number]::-webkit-outer-spin-button,
.form-input[type=number]::-webkit-inner-spin-button {
    -webkit-appearance: none;
    margin: 0;
}

/* Firefox */
/* ========== SELECT WRAPPER ========== */
.select-wrapper {
    position: relative;
}

.form-select {
    width: 100%;
    padding: 14px 16px;
    padding-right: 45px;
    border: 2px solid #e5e5e5;
    border-radius: 10px;
    font-size: 15px;
    color: #333;
    transition: all 0.2s;
    font-family: "Kanit", sans-serif;
    appearance: none;
    background: #fff;
    cursor: pointer;
}

.form-select:focus {
    outline: none;
    border-color: #037266;
    background: #f0f8f7;
}

.select-icon {
    position: absolute;
    right: 16px;
    top: 50%;
    transform: translateY(-50%);
    pointer-events: none;
    color: #666;
}

/* ========== PERIOD DROPDOWN ========== */
.period-header {
    display: flex;
    align-items: center;
    gap: 12px;
    margin-bottom: 14px;
}

.period-header label {
    display: flex;
    align-items: center;
    font-size: 15px;
    font-weight: 600;
    color: #333;
    margin: 0;
}

.period-dropdown-wrapper {
    display: flex;
    align-items: flex-end;
    gap: 12px;
}

.period-select-container {
    position: relative;
    flex: 1;
    min-width: 200px;
}

.period-select {
    width: 100%;
    padding: 12px 16px;
    padding-right: 40px;
    border: 2px solid #e5e5e5;
    border-radius: 10px;
    font-size: 15px;
    font-weight: 500;
    color: #333;
    background: #fff;
    cursor: pointer;
    transition: all 0.2s ease;
    font-family: "Kanit", sans-serif;
    appearance: none;
}

.period-select:hover {
    border-color: #cee5e2;
    background: #f9fffe;
}

.period-select:focus {
    outline: none;
    border-color: #037266;
    background: #f0f8f7;
    box-shadow: 0 0 0 3px rgba(3, 114, 102, 0.1);
}

.select-chevron {
    position: absolute;
    right: 14px;
    top: 50%;
    transform: translateY(-50%);
    pointer-events: none;
    color: #666;
    stroke-width: 2.5;
}

.custom-input-wrapper {
    display: flex;
    align-items: center;
    gap: 10px;
}

.custom-days-input {
    width: 140px;
    padding: 12px 14px;
    border: 2px solid #e5e5e5;
    border-radius: 10px;
    font-size: 15px;
    color: #333;
    background: #fff;
    transition: all 0.2s ease;
    font-family: "Kanit", sans-serif;
    font-weight: 500;
}

.custom-days-input:hover {
    border-color: #cee5e2;
    background: #f9fffe;
}

.custom-days-input:focus {
    outline: none;
    border-color: #037266;
    background: #f0f8f7;
    box-shadow: 0 0 0 3px rgba(3, 114, 102, 0.1);
}

.custom-days-input::placeholder {
    color: #999;
}

.days-suffix {
    font-size: 15px;
    font-weight: 500;
    color: #666;
    white-space: nowrap;
}

/* ========== TOOLTIP ========== */
.tooltip-wrapper {
    position: relative;
}

.tooltip-btn {
    width: 28px;
    height: 28px;
    border-radius: 50%;
    background: #f0f0f0;
    border: 2px solid #e0e0e0;
    color: #666;
    cursor: pointer;
    display: flex;
    align-items: center;
    justify-content: center;
    transition: all 0.2s ease;
    padding: 0;
    flex-shrink: 0;
}

.tooltip-btn:hover {
    background: #037266;
    border-color: #037266;
    color: #fff;
}

.tooltip-content {
    position: absolute;
    top: 100%;
    right: 0;
    background: #fffbf0;
    border: 2px solid #ffd43b;
    border-radius: 12px;
    padding: 16px;
    margin-top: 8px;
    width: 280px;
    box-shadow: 0 4px 16px rgba(0, 0, 0, 0.1);
    z-index: 1000;
}

.tooltip-content::before {
    content: '';
    position: absolute;
    bottom: 100%;
    right: 12px;
    width: 10px;
    height: 10px;
    background: #fffbf0;
    border: 2px solid #ffd43b;
    border-bottom: none;
    border-left: none;
    transform: rotate(45deg);
}

.tooltip-content p {
    font-size: 13px;
    color: #666;
    margin: 0 0 8px 0;
    line-height: 1.5;
}

.tooltip-content p:first-child {
    color: #333;
    font-weight: 600;
    margin-bottom: 10px;
}

.tooltip-content ul {
    list-style: none;
    padding: 0;
    margin: 0 0 10px 0;
}

.tooltip-content li {
    font-size: 13px;
    color: #666;
    margin-bottom: 6px;
    line-height: 1.5;
    padding-left: 16px;
    position: relative;
}

.tooltip-content li:before {
    content: '•';
    position: absolute;
    left: 0;
    color: #f59e0b;
}

.tooltip-content li:last-child {
    margin-bottom: 0;
}

.notice-highlight {
    color: #f59e0b !important;
    font-weight: 600 !important;
    margin-top: 10px;
}

/* ========== EDITOR TOOLBAR ========== */
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

.toolbar-btn:active {
    background: #d0d0d0;
}

.toolbar-divider {
    width: 1px;
    height: 24px;
    background: #d0d0d0;
    margin: 0 4px;
}

/* ========== FORM EDITOR ========== */
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

/* ========== CHECKBOX GROUP ========== */
.checkbox-group {
    margin-bottom: 15px;
}

.checkbox-item {
    display: block;
    cursor: pointer;
}

.checkbox-item input[type="checkbox"] {
    display: none;
}

.checkbox-content {
    display: flex;
    align-items: center;
    gap: 12px;
    padding: 14px 16px;
    background: #fff;
    border: 2px solid #e5e5e5;
    border-radius: 10px;
    transition: all 0.2s;
    font-size: 15px;
    font-weight: 500;
    color: #333;
}

.checkbox-item:hover .checkbox-content {
    border-color: #cee5e2;
    background: #f0f8f7;
}

.checkbox-item input:checked+.checkbox-content {
    border-color: #037266;
    background: #f0f8f7;
}

.checkbox-box {
    width: 22px;
    height: 22px;
    border: 2px solid #d0d0d0;
    border-radius: 6px;
    position: relative;
    flex-shrink: 0;
    transition: all 0.2s;
}

.checkbox-item input:checked+.checkbox-content .checkbox-box {
    border-color: #037266;
    background: #037266;
}

.checkbox-item input:checked+.checkbox-content .checkbox-box::after {
    content: '';
    position: absolute;
    top: 3px;
    left: 7px;
    width: 5px;
    height: 10px;
    border: solid #fff;
    border-width: 0 2px 2px 0;
    transform: rotate(45deg);
}

/* ========== PREVIEW CARD ========== */
.preview-card {
    background: #f9f9f9;
    border-radius: 16px;
    padding: 30px;
    box-shadow: 0 2px 12px rgba(3, 114, 102, 0.07);
    border: 1px solid #eef0ef;
    margin-bottom: 25px;
}

.preview-card h3 {
    font-size: 20px;
    font-weight: 700;
    color: #222;
    margin: 0 0 10px 0;
}

/* ========== JOB PREVIEW ========== */
.job-preview {
    background: #fff;
    border-radius: 12px;
    overflow: hidden;
    margin-top: 20px;
    box-shadow: 0 2px 8px rgba(0, 0, 0, 0.08);
}

.preview-images {
    position: relative;
    width: 100%;
    height: 300px;
    background: #f0f0f0;
}

.preview-main-image {
    width: 100%;
    height: 100%;
    object-fit: cover;
}

.preview-image-count {
    position: absolute;
    bottom: 15px;
    right: 15px;
    background: rgba(0, 0, 0, 0.7);
    color: #fff;
    padding: 8px 14px;
    border-radius: 20px;
    display: flex;
    align-items: center;
    gap: 6px;
    font-size: 13px;
    font-weight: 600;
    backdrop-filter: blur(4px);
}

.preview-content {
    padding: 30px;
}

.preview-title {
    font-size: 24px;
    font-weight: 700;
    color: #222;
    margin: 0 0 10px 0;
}

.preview-company {
    font-size: 15px;
    color: #666;
    margin: 0 0 20px 0;
}

.preview-meta {
    display: flex;
    flex-wrap: wrap;
    gap: 20px;
    margin-bottom: 25px;
    padding-bottom: 20px;
    border-bottom: 2px solid #f0f0f0;
}

.meta-item {
    display: flex;
    align-items: center;
    gap: 8px;
    font-size: 14px;
    color: #555;
}

.meta-item svg {
    color: #037266;
    flex-shrink: 0;
}

.preview-description h4 {
    font-size: 18px;
    font-weight: 600;
    color: #222;
    margin: 0 0 15px 0;
}

.preview-description {
    font-size: 15px;
    color: #555;
    line-height: 1.8;
}

/* ========== CONFIRMATION CARD ========== */
.confirmation-card {
    background: #f9f9f9;
    border-radius: 16px;
    padding: 40px;
    box-shadow: 0 2px 12px rgba(3, 114, 102, 0.07);
    border: 1px solid #eef0ef;
    text-align: center;
}

.confirmation-icon {
    width: 80px;
    height: 80px;
    background: linear-gradient(135deg, #037266 0%, #026b5b 100%);
    border-radius: 50%;
    display: flex;
    align-items: center;
    justify-content: center;
    margin: 0 auto 25px;
    color: #fff;
    box-shadow: 0 4px 16px rgba(3, 114, 102, 0.2);
}

.confirmation-card h2 {
    font-size: 24px;
    font-weight: 700;
    color: #222;
    margin: 0 0 15px 0;
}

.confirmation-text {
    font-size: 16px;
    color: #555;
    line-height: 1.7;
    margin-bottom: 30px;
}

.confirmation-text strong {
    color: #037266;
    font-weight: 700;
}

/* ========== CONFIRMATION DETAILS ========== */
.confirmation-details {
    background: #fff;
    border-radius: 12px;
    padding: 25px;
    margin-bottom: 25px;
    text-align: left;
}

.detail-item {
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 14px 0;
    border-bottom: 1px solid #f0f0f0;
    gap: 20px;
}

.detail-item:last-child {
    border-bottom: none;
}

.detail-label {
    font-size: 15px;
    color: #666;
    font-weight: 500;
}

.detail-value {
    font-size: 15px;
    color: #333;
    font-weight: 600;
    text-align: right;
    word-break: break-word;
}

.confirmation-notice {
    background: #f0f9ff;
    border: 2px solid #bae6fd;
    border-radius: 12px;
    padding: 18px 20px;
    display: flex;
    align-items: flex-start;
    gap: 12px;
    text-align: left;
}

.confirmation-notice svg {
    color: #0369a1;
    flex-shrink: 0;
    margin-top: 2px;
}

.confirmation-notice p {
    font-size: 14px;
    color: #0c4a6e;
    margin: 0;
    line-height: 1.6;
    font-weight: 500;
}

/* ========== ACTION BUTTONS ========== */
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

.btn-confirm {
    background: linear-gradient(135deg, #10b981 0%, #059669 100%);
}

.btn-confirm:hover {
    background: linear-gradient(135deg, #059669 0%, #047857 100%);
}

/* ========== STEP CONTENT ANIMATION ========== */
.step-content {
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

/* ========== RESPONSIVE ========== */
@media (max-width: 768px) {
    .create-job-page {
        padding: 15px;
    }

    .page-header h1 {
        font-size: 24px;
    }

    .progress-steps {
        padding: 0;
    }

    .step-label {
        font-size: 12px;
    }

    .step-number {
        width: 40px;
        height: 40px;
        font-size: 16px;
    }

    .step-line {
        margin: 0 10px;
    }

    .section-card,
    .preview-card,
    .confirmation-card {
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
    }

    .detail-item {
        flex-direction: column;
        align-items: flex-start;
        gap: 5px;
    }

    .detail-value {
        text-align: left;
    }

    .confirmation-icon {
        width: 70px;
        height: 70px;
    }

    .period-dropdown-wrapper {
        flex-wrap: wrap;
    }

    .period-select-container {
        width: 100%;
    }

    .custom-input-wrapper {
        width: 100%;
    }

    .custom-days-input {
        flex: 1;
    }

    .tooltip-content {
        right: auto;
        left: 0;
    }

    .tooltip-content::before {
        right: auto;
        left: 12px;
    }

    .period-header {
        flex-wrap: wrap;
        gap: 8px;
    }
}

@media (max-width: 480px) {
    .page-header h1 {
        font-size: 22px;
    }

    .progress-steps {
        flex-wrap: wrap;
        gap: 20px;
    }

    .step-line {
        display: none;
    }

    .step {
        flex: 1 1 calc(33.333% - 15px);
        min-width: 80px;
    }

    .uploaded-images {
        grid-template-columns: repeat(2, 1fr);
    }

    .preview-title {
        font-size: 20px;
    }

    .confirmation-card {
        padding: 30px 20px;
    }

    .confirmation-card h2 {
        font-size: 20px;
    }

    .period-dropdown-wrapper {
        flex-direction: column;
        width: 100%;
    }

    .period-select-container {
        width: 100%;
    }

    .custom-input-wrapper {
        width: 100%;
    }

    .custom-days-input {
        width: 100%;
    }
}
</style>