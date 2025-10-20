<!-- frontend/src/pages/Company/CompanyCreateJobs.vue -->
<template>
    <DashboardLayout role="company">
        <div class="create-job-page">
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

            <div class="progress-steps">
                <div class="step" :class="{ active: currentStep >= 1, completed: currentStep > 1 }">
                    <div class="step-number">1</div>
                    <div class="step-label">สร้างประกาศ</div>
                </div>
                <div class="step-line" :class="{ active: currentStep > 1 }"></div>
                <div class="step" :class="{ active: currentStep >= 2, completed: currentStep > 2 }">
                    <div class="step-number">2</div>
                    <div class="step-label">ตรวจสอบและยืนยันการประกาศ</div>
                </div>
                <div class="step-line" :class="{ active: currentStep > 2 }"></div>
                <div class="step" :class="{ active: currentStep >= 3 }">
                    <div class="step-number">3</div>
                    <div class="step-label">ผลการประกาศ</div>
                </div>
            </div>

            <div v-if="currentStep === 1" class="step-content">
                <div class="section-card">
                    <h3>รูปโปสเตอร์</h3>
                    <p class="helper-text">อัปโหลดรูปภาพประกอบเพื่อดึงดูดความสนใจจากผู้สมัครงาน (ไม่เกิน 5 รูป)</p>

                    <div class="upload-area">
<div class="uploaded-images" v-if="uploadedImages.length > 0">
    <div v-for="(image, index) in uploadedImages" :key="index" class="image-preview-item"
        :class="{ active: currentImageIndex === index }"
        @click="currentImageIndex = index"
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
                            <p>คลิกเพื่ออัปโหลดรูปภาพ</p>
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
                    <h3>เขียนประกาศงาน <span class="required">*</span></h3>
                    <p class="helper-text">เขียนรายละเอียดเกี่ยวกับตำแหน่งงานและคุณสมบัติผู้สมัคร</p>

                    <div class="form-group">
                        <label>ตำแหน่งงาน </label>
                        <input type="text" v-model="jobForm.title" placeholder="เช่น Frontend Developer Intern"
                            class="form-input" />
                    </div>

                    <div class="form-group">
                        <label>ภาควิชา</label>
                        <div class="custom-dropdown" ref="deptDropdownRef"
                            @click="isDeptDropdownOpen = !isDeptDropdownOpen">
                            <div class="dropdown-selected">
                                <span>{{ jobForm.department || 'เลือกภาควิชา' }}</span>
                                <svg class="dropdown-arrow" :class="{ open: isDeptDropdownOpen }"
                                    xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24"
                                    fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round"
                                    stroke-linejoin="round">
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
                        <p class="field-hint">เลือกภาควิชาที่ต้องการกำหนดให้ผู้สมัครงาน</p>
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
                            placeholder="เขียนรายละเอียดงาน คุณสมบัติผู้สมัคร และสิทธิประโยชน์..."></div>
                    </div>
                </div>

                <div class="section-card">
                    <h3>กำหนดระยะเวลารับสมัครงาน <span class="required">*</span></h3>
                    <p class="helper-text">กำหนดช่วงเวลาที่เปิดรับสมัครงานและระยะเวลาการพิจารณา</p>

                    <div class="form-group">
                        <div class="period-header">
                            <label>ระยะเวลารับสมัครงาน</label>
                            <div class="tooltip-wrapper">
                                <button type="button" class="tooltip-btn" @mouseenter="showTooltip = true"
                                    @mouseleave="showTooltip = false">
                                    <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24"
                                        fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round"
                                        stroke-linejoin="round">
                                        <circle cx="12" cy="12" r="10"></circle>
                                        <line x1="12" y1="8" x2="12" y2="12"></line>
                                        <line x1="12" y1="16" x2="12.01" y2="16"></line>
                                    </svg>
                                </button>
                                <div v-if="showTooltip" class="tooltip-content">
                                    <p><strong>ตัวอย่างการกำหนดระยะเวลารับสมัคร</strong></p>
                                    <p>
                                        หากประกาศงานได้รับการอนุมัติในวันที่ <strong>1 ม.ค. 2568</strong>
                                        และตั้งค่า “ระยะเวลารับสมัคร <strong>15 วัน</strong>”
                                        ระบบจะเริ่มเผยแพร่ประกาศในวันที่ได้รับการอนุมัติ
                                        และกำหนดวันปิดรับสมัครเป็นวันที่ <strong>16 ม.ค. 2568</strong>
                                    </p>
                                    <p class="notice-highlight">
                                        <strong>หมายเหตุ:</strong> ระบบจะปิดรับสมัครให้อัตโนมัติตามวันที่กำหนดไว้
                                    </p>
                                </div>
                            </div>
                        </div>

                        <div class="period-dropdown-wrapper">
                            <div class="custom-dropdown period-select-container" ref="durationDropdownRef"
                                @click="isDurationDropdownOpen = !isDurationDropdownOpen">
                                <div class="dropdown-selected">
                                    <span>{{DATE_OPTIONS.find(opt => opt.value === selectedDateOption)?.label}}</span>
                                    <svg class="dropdown-arrow" :class="{ open: isDurationDropdownOpen }"
                                        xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24"
                                        fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round"
                                        stroke-linejoin="round">
                                        <polyline points="6 9 12 15 18 9"></polyline>
                                    </svg>
                                </div>
                                <transition name="dropdown-fade">
                                    <ul v-if="isDurationDropdownOpen" class="dropdown-menu">
                                        <li v-for="option in DATE_OPTIONS" :key="option.value"
                                            @click.stop="selectDurationOption(option)">
                                            {{ option.label }}
                                        </li>
                                    </ul>
                                </transition>
                            </div>

                            <div v-if="selectedDateOption === 'custom'" class="custom-input-wrapper">
                                <input type="number" v-model.number="customDays" placeholder="กรอกจำนวนวัน"
                                    class="custom-days-input" min="1" />
                                <span class="days-suffix">วัน</span>
                            </div>
                        </div>
                    </div>
                </div>

                <div class="section-card">
                    <h3>วิธีรับข้อมูลการสมัคร</h3>
                    <p class="helper-text">ระบุอีเมลที่ต้องการรับการแจ้งเตือนเมื่อมีผู้สมัครงานผ่านระบบ Science
                        Cooperative Education System</p>

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
                    </div>
                </div>

                <div class="action-buttons">
                    <button class="btn-secondary" @click="cancelCreation">ยกเลิก</button>
                    <button class="btn-primary" @click="nextStep" :disabled="!isStep1Valid">ถัดไป</button>
                </div>
            </div>

            <div v-if="currentStep === 2" class="step-content">
                <div class="preview-tabs">
                    <button class="tab-btn" :class="{ active: previewMode === 'card' }" @click="previewMode = 'card'">
                        <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none"
                            stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                            <rect x="3" y="3" width="7" height="7"></rect>
                            <rect x="14" y="3" width="7" height="7"></rect>
                            <rect x="14" y="14" width="7" height="7"></rect>
                            <rect x="3" y="14" width="7" height="7"></rect>
                        </svg>
                        <span>มุมมองการ์ด</span>
                    </button>
                    <button class="tab-btn" :class="{ active: previewMode === 'panel' }" @click="previewMode = 'panel'">
                        <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none"
                            stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                            <rect x="3" y="3" width="18" height="18" rx="2" ry="2"></rect>
                            <line x1="9" y1="3" x2="9" y2="21"></line>
                        </svg>
                        <span>มุมมองแบบเต็ม</span>
                    </button>
                </div>
                <div v-if="previewMode === 'card'" class="preview-container">

                    <div class="job-card-preview">
                        <div class="job-header">
                            <img class="company-logo" :src="mockCompanyData.logo" alt="logo" />
                            <div class="job-header-info">
                                <div class="company-name">{{ mockCompanyData.name }}</div>
                                <div class="department-label">{{ jobForm.department || 'ไม่ระบุภาควิชา' }}</div>
                                <div class="job-meta">
                                    <div class="job-meta-row">
                                        <span class="meta-icon">
                                            <svg width="16" height="16" fill="none" stroke="#9CA3AF" stroke-width="2"
                                                stroke-linecap="round" stroke-linejoin="round" viewBox="0 0 24 24">
                                                <rect x="3" y="4" width="18" height="18" rx="2" ry="2" />
                                                <line x1="16" y1="2" x2="16" y2="6" />
                                                <line x1="8" y1="2" x2="8" y2="6" />
                                                <line x1="3" y1="10" x2="21" y2="10" />
                                            </svg>
                                        </span>
                                        <span>ประกาศเมื่อ: รออนุมัติจากอาจารย์ผู้ประสานงานที่เกี่ยวข้อง</span>
                                    </div>
                                    <div class="job-meta-row">
                                        <span class="meta-icon">
                                            <svg width="16" height="16" fill="none" stroke="#9CA3AF" stroke-width="2"
                                                stroke-linecap="round" stroke-linejoin="round" viewBox="0 0 24 24">
                                                <circle cx="12" cy="12" r="10"></circle>
                                                <polyline points="12 6 12 12 16 14"></polyline>
                                            </svg>
                                        </span>
                                        <span>ปิดรับสมัคร: {{ getDateRangePreview() }}</span>
                                    </div>
                                </div>
                            </div>
                        </div>

                        <div class="job-content">
                            <div class="job-images" v-if="uploadedImages.length > 0">
                                <img v-if="uploadedImages.length === 1" :src="uploadedImages[0].preview" class="job-img"
                                    alt="job" />
                                <div v-else class="image-grid" :class="`images-${Math.min(uploadedImages.length, 4)}`">
                                    <div v-for="(img, i) in uploadedImages.slice(0, 4)" :key="i" class="image-cell"
                                        @click="currentImageIndex = i" style="cursor: pointer;">
                                        <img :src="img.preview" alt="job" />
                                        <div v-if="i === 3 && uploadedImages.length > 4" class="image-overlay">
                                            +{{ uploadedImages.length - 4 }}
                                        </div>
                                    </div>
                                </div>
                            </div>

                            <div class="job-details">
                                <div class="job-title">{{ jobForm.title || 'ตำแหน่งงาน' }}</div>
                                <div class="job-details-title">รายละเอียดงาน</div>
                                <div class="job-description" v-html="truncateHTML(jobForm.description, 160)"></div>
                                <span v-if="jobForm.description.length > 160" class="show-more">... ดูเพิ่มเติม</span>
                            </div>
                        </div>
                    </div>
                </div>

                <div v-if="previewMode === 'panel'" class="preview-container">

                    <div class="side-panel-preview">
                        <div class="side-panel-inner">
                            <div class="side-header-top">
                                <img class="side-logo" :src="mockCompanyData.logo" alt="logo" />
                                <div class="side-title-block">
                                    <div class="side-title">{{ jobForm.title || 'ตำแหน่งงาน' }}</div>
                                    <div class="side-company-row">
                                        <span class="side-company-name">{{ mockCompanyData.name }}</span>
                                        <span class="side-company-sep"></span>
                                        <span class="side-company-all-jobs">ดูงานทั้งหมด</span>
                                    </div>
                                    <span class="side-department-label">
                                        {{ jobForm.department || 'ไม่ระบุภาควิชา' }}
                                    </span>
                                </div>
                            </div>

                            <div class="side-meta-row">
                                <span class="side-meta-item">
                                    <svg width="22" height="19" fill="none" stroke="#9CA3AF" stroke-width="2"
                                        stroke-linecap="round" stroke-linejoin="round" viewBox="0 0 24 24">
                                        <rect x="3" y="4" width="18" height="18" rx="2" ry="2" />
                                        <line x1="16" y1="2" x2="16" y2="6" />
                                        <line x1="8" y1="2" x2="8" y2="6" />
                                        <line x1="3" y1="10" x2="21" y2="10" />
                                    </svg>
                                    ประกาศเมื่อ: รออนุมัติจากอาจารย์ผู้ประสานงานที่เกี่ยวข้อง
                                </span>
                            </div>

                            <div class="side-meta-row">
                                <span class="side-meta-item">
                                    <svg width="22" height="19" fill="none" stroke="#9CA3AF" stroke-width="2"
                                        stroke-linecap="round" stroke-linejoin="round" viewBox="0 0 24 24">
                                        <circle cx="12" cy="12" r="10"></circle>
                                        <polyline points="12 6 12 12 16 14"></polyline>
                                    </svg>
                                    ปิดรับสมัคร: {{ getDateRangePreview() }}
                                </span>
                            </div>

                            <div class="side-meta-row">
                                <span class="side-meta-item">
                                    <svg width="22" height="22" fill="none" stroke="#9CA3AF" stroke-width="2"
                                        stroke-linecap="round" stroke-linejoin="round" viewBox="0 0 24 24">
                                        <path d="M22 10L12 4 2 10l10 6 10-6z"></path>
                                        <path d="M6 12v5c0 1.1 2.7 2 6 2s6-.9 6-2v-5"></path>
                                    </svg>
                                    ปีการศึกษา {{ mockCompanyData.academicYear }} ภาคที่ {{ mockCompanyData.semester }}
                                </span>
                            </div>

                            <div v-if="uploadedImages.length > 0" class="side-images-preview">
                                <div class="side-image-layout">
                                    <div class="side-image-large">
                                        <img :src="uploadedImages[currentImageIndex].preview" alt="job-image-large" />
                                    </div>
                                    <div class="side-image-thumbs-wrapper" v-if="uploadedImages.length > 1">
                                        <div class="side-image-thumbs">
                                            <div v-for="(img, i) in uploadedImages" :key="i" class="side-image-thumb"
                                                :class="{ active: currentImageIndex === i }"
                                                @click="currentImageIndex = i">
                                                <img :src="img.preview" alt="job-thumb" />
                                            </div>
                                        </div>
                                    </div>
                                </div>
                            </div>

                            <div class="side-detail-content">
                                <div class="side-detail-title">รายละเอียดงาน</div>
                                <div class="side-description" v-html="jobForm.description || 'ไม่มีรายละเอียด'"></div>
                            </div>
                        </div>
                    </div>
                </div>

                <div v-if="!hasPreviewContent" class="empty-preview-state">
                    <svg xmlns="http://www.w3.org/2000/svg" width="64" height="64" viewBox="0 0 24 24" fill="none"
                        stroke="currentColor" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round">
                        <path d="M14 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8z"></path>
                        <polyline points="14 2 14 8 20 8"></polyline>
                        <line x1="12" y1="18" x2="12" y2="12"></line>
                        <line x1="9" y1="15" x2="15" y2="15"></line>
                    </svg>
                    <h3>ยังไม่มีข้อมูลสำหรับแสดงตัวอย่าง</h3>
                    <p>กรุณากลับไปขั้นตอนที่ 1 เพื่อกรอกข้อมูลประกาศงาน</p>
                    <button class="btn-secondary" @click="prevStep">
                        <svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none"
                            stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                            <line x1="19" y1="12" x2="5" y2="12"></line>
                            <polyline points="12 19 5 12 12 5"></polyline>
                        </svg>
                        กลับไปขั้นตอนที่ 1
                    </button>
                </div>

                <div class="section-card confirmation-card">
                    <h3 style="text-align: center;">ยืนยันการประกาศงาน</h3>
                    <p class="confirmation-text">
                        คุณต้องการยืนยันการประกาศงานตำแหน่ง
                        <strong>{{ jobForm.title }}</strong> ใช่หรือไม่?
                    </p>

                    <div class="confirmation-notice">
                        <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none"
                            stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                            <circle cx="12" cy="12" r="10"></circle>
                            <line x1="12" y1="8" x2="12" y2="12"></line>
                            <line x1="12" y1="16" x2="12.01" y2="16"></line>
                        </svg>
                        <p>เมื่อกดยืนยันแล้ว
                            ประกาศงานของคุณจะถูกส่งไปยังอาจารย์ผู้ประสานงานภาควิชาเพื่อตรวจสอบก่อนเผยเเพร่</p>
                    </div>
                </div>

                <div class="action-buttons">
                    <button class="btn-secondary" @click="prevStep">ย้อนกลับ</button>
                    <button class="btn-primary" @click="nextStep">ถัดไป</button>
                </div>
            </div>

            <div v-if="currentStep === 3" class="step-content">
                <div class="application-summary-card">
                    <!-- Header with Icon -->
                    <div class="summary-header">
                        <div class="summary-icon">
                            <svg xmlns="http://www.w3.org/2000/svg" width="48" height="48" viewBox="0 0 24 24"
                                fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round"
                                stroke-linejoin="round">
                                <path d="M22 11.08V12a10 10 0 1 1-5.93-9.14"></path>
                                <polyline points="22 4 12 14.01 9 11.01"></polyline>
                            </svg>
                        </div>
                        <h2>ประกาศงานสำเร็จเรียบร้อยแล้ว</h2>
                        <p class="summary-subtitle">กรุณาตรวจสอบรายละเอียดประกาศงานของคุณ</p>
                    </div>

                    <!-- Job Position Section -->
                    <div class="summary-section">
                        <div class="section-header-summary">
                            <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24"
                                fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round"
                                stroke-linejoin="round">
                                <rect x="2" y="7" width="20" height="14" rx="2" ry="2"></rect>
                                <path d="M16 21V5a2 2 0 0 0-2-2h-4a2 2 0 0 0-2 2v16"></path>
                            </svg>
                            <h4>ตำแหน่งงานที่ประกาศ</h4>
                        </div>
                        <div class="summary-content">
                            <div class="job-position-display">
                                <img :src="mockCompanyData.logo" alt="Company Logo" class="job-logo-small" />
                                <div class="job-details-summary">
                                    <h5>{{ jobForm.title }}</h5>
                                    <p class="company-detail">{{ mockCompanyData.name }}</p>
                                    <span class="department-badge">{{ jobForm.department || 'ไม่ระบุภาควิชา' }}</span>
                                </div>
                            </div>
                        </div>
                    </div>

                    <!-- Job Details Section -->
                    <div class="summary-section">
                        <div class="section-header-summary">
                            <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24"
                                fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round"
                                stroke-linejoin="round">
                                <path d="M14 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8z"></path>
                                <polyline points="14 2 14 8 20 8"></polyline>
                            </svg>
                            <h4>รายละเอียดประกาศ</h4>
                        </div>
                        <div class="summary-content">
                            <div class="applicant-info-grid">
                                <div class="info-row-summary">
                                    <span class="info-label-summary">ตำแหน่ง</span>
                                    <span class="info-value-summary">{{ jobForm.title }}</span>
                                </div>
                                <div class="info-row-summary">
                                    <span class="info-label-summary">ภาควิชา</span>
                                    <span class="info-value-summary">{{ jobForm.department || 'ไม่ระบุภาควิชา' }}</span>
                                </div>
                                <div class="info-row-summary">
                                    <span class="info-label-summary">ระยะเวลารับสมัคร</span>
                                    <span class="info-value-summary">{{ getDateRangePreview() }}</span>
                                </div>
                                <div class="info-row-summary">
                                    <span class="info-label-summary">อีเมลรับการแจ้งเตือน</span>
                                    <span class="info-value-summary">
                                        <svg xmlns="http://www.w3.org/2000/svg" width="14" height="14"
                                            viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"
                                            stroke-linecap="round" stroke-linejoin="round">
                                            <path
                                                d="M4 4h16c1.1 0 2 .9 2 2v12c0 1.1-.9 2-2 2H4c-1.1 0-2-.9-2-2V6c0-1.1.9-2 2-2z">
                                            </path>
                                            <polyline points="22,6 12,13 2,6"></polyline>
                                        </svg>
                                        {{ jobForm.notificationEmail || 'ไม่มีอีเมล' }}
                                    </span>
                                </div>
                            </div>
                        </div>
                    </div>

                    <!-- Images Section (if any) -->
                    <div v-if="uploadedImages.length > 0" class="summary-section">
                        <div class="section-header-summary">
                            <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24"
                                fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round"
                                stroke-linejoin="round">
                                <rect x="3" y="3" width="18" height="18" rx="2" ry="2"></rect>
                                <circle cx="8.5" cy="8.5" r="1.5"></circle>
                                <polyline points="21 15 16 10 5 21"></polyline>
                            </svg>
                            <h4>รูปภาพประกอบ</h4>
                        </div>
                        <div class="summary-content">
                            <div class="documents-list">
                                <div class="document-item">
                                    <div class="document-icon resume-icon">
                                        <svg xmlns="http://www.w3.org/2000/svg" width="18" height="18"
                                            viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"
                                            stroke-linecap="round" stroke-linejoin="round">
                                            <rect x="3" y="3" width="18" height="18" rx="2" ry="2"></rect>
                                            <circle cx="8.5" cy="8.5" r="1.5"></circle>
                                            <polyline points="21 15 16 10 5 21"></polyline>
                                        </svg>
                                    </div>
                                    <div class="document-info">
                                        <span class="document-type">รูปภาพประกอบ</span>
                                        <span class="document-name">จำนวน {{ uploadedImages.length }} รูป</span>
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>

                    <div class="application-date">
                        <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none"
                            stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                            <rect x="3" y="4" width="18" height="18" rx="2" ry="2"></rect>
                            <line x1="16" y1="2" x2="16" y2="6"></line>
                            <line x1="8" y1="2" x2="8" y2="6"></line>
                            <line x1="3" y1="10" x2="21" y2="10"></line>
                        </svg>
                        <span>ส่งประกาศเมื่อวันที่ {{ formatDate(new Date()) }}</span>
                    </div>
                </div>

                <!-- Job Status Info -->
                <!-- Job Status Timeline -->
                <div class="status-timeline-card">
                    <div class="timeline-header">
                        <h3>ขั้นตอนการประกาศงาน</h3>
                    </div>

                    <div class="timeline-container">
                        <div class="timeline-wrapper">
                            <div v-for="(stage, index) in jobStages" :key="index" class="timeline-item">
                                <div class="timeline-dot" :class="stage.status">
                                    <svg v-if="stage.status === 'passed'" xmlns="http://www.w3.org/2000/svg" width="20"
                                        height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor"
                                        stroke-width="3" stroke-linecap="round" stroke-linejoin="round">
                                        <polyline points="20 6 9 17 4 12"></polyline>
                                    </svg>
                                    <svg v-else-if="stage.status === 'current'" xmlns="http://www.w3.org/2000/svg"
                                        width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor"
                                        stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                                        <circle cx="12" cy="12" r="10"></circle>
                                        <line x1="12" y1="8" x2="12" y2="12"></line>
                                        <line x1="12" y1="16" x2="12.01" y2="16"></line>
                                    </svg>
                                </div>
                                <div class="timeline-content">
                                    <div class="timeline-label">{{ stage.name }}</div>
                                    <p class="timeline-description">{{ stage.description }}</p>
                                </div>
                                <div v-if="index < jobStages.length - 1" class="timeline-connector"
                                    :class="getConnectorClass(index)"></div>
                            </div>
                        </div>
                    </div>

                    <div class="timeline-info">
                        <div class="info-box">
                            <svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24"
                                fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round"
                                stroke-linejoin="round">
                                <circle cx="12" cy="12" r="10"></circle>
                                <line x1="12" y1="8" x2="12" y2="12"></line>
                                <line x1="12" y1="16" x2="12.01" y2="16"></line>
                            </svg>
                            <p>ระบบจะแจ้งเตือนผ่านอีเมลเมื่อมีนักศึกษาสมัครงาน
                                คุณสามารถตรวจสอบและจัดการผู้สมัครได้ที่หน้าจัดการงาน</p>
                        </div>
                    </div>
                </div>

                <div class="action-buttons">
                    <button class="btn-secondary"
                        @click="router.push({ name: 'CompanyJobs' })">กลับสู่หน้าจัดการงาน</button>
                    <button class="btn-primary" @click="router.push({ name: 'CompanyJobs' })">ดูประกาศงาน</button>
                </div>
            </div>
        </div>
    </DashboardLayout>
</template>

<script setup>
import { ref, reactive, computed, watch, onMounted, onBeforeUnmount } from 'vue';
import { useRouter } from 'vue-router';
import DashboardLayout from '../../components/DashboardLayout.vue';

// =======================================================
// 1. CONSTANTS & STATIC DATA
// =======================================================
const DATE_OPTIONS = [
    { value: '7', label: '7 วัน', days: 7 },
    { value: '15', label: '15 วัน', days: 15 },
    { value: '20', label: '20 วัน', days: 20 },
    { value: 'custom', label: 'กำหนดเอง', days: null }
];

const departments = [
    "ภาควิชาคณิตศาสตร์",
    "ภาควิชาเคมี",
    "ภาควิชาชีววิทยา",
    "ภาควิชาฟิสิกส์",
    "ภาควิชาวิทยาศาสตร์สิ่งแวดล้อม",
    "ภาควิชาสถิติ",
    "ภาควิชาคอมพิวเตอร์",
    "ภาควิชาจุลชีววิทยา"
];

// Mock company data (ข้อมูลชั่วคราวสำหรับพรีวิว)
const mockCompanyData = {
    name: 'บริษัท XYZ จำกัด',
    logo: 'https://via.placeholder.com/120', // หรือใช้โลโก้จริงจาก auth/profile
    academicYear: 2568,
    semester: 1
};

// Job Stages for Timeline (Company side)
const jobStages = ref([
    {
        name: "สร้างประกาศงาน",
        status: "passed",
        description: "สร้างประกาศงานเบื้องต้นและส่งเรียบร้อยแล้ว"
    },
    {
        name: "รอการอนุมัติ",
        status: "current",
        description: "ประกาศงานกำลังรอการพิจารณาอนุมัติจากอาจารย์ผู้ประสานงานภาควิชา"
    },
    {
        name: "เผยแพร่ประกาศ",
        status: "pending",
        description: "เมื่อได้รับอนุมัติ ประกาศจะถูกเผยแพร่และเริ่มนับระยะเวลารับสมัครทันที"
    }
]);


// =======================================================
// 2. STATE MANAGEMENT (REFS & REACTIVES)
// =======================================================
const router = useRouter();

// A. Navigation & Flow State
const currentStep = ref(1);
const previewMode = ref('card'); // 'card' หรือ 'panel'

// B. Form Data (Reactive)
const jobForm = reactive({
    title: '',
    department: '',
    description: '',
    sendToCompanyEmail: false,
    notificationEmail: '',
    applicationDays: 15 // Stores the actual number of days
});
// Set initial applicationDays based on default selectedDateOption ('15')
jobForm.applicationDays = DATE_OPTIONS.find(o => o.value === '15').days;

// C. UI & Dropdown State
const showPreviewNotice = ref(true);
const uploadedImages = ref([]);
const currentImageIndex = ref(0);
const showTooltip = ref(false);
const customDays = ref(1);
const selectedDateOption = ref('15'); // Default selection
const isDeptDropdownOpen = ref(false);
const isDurationDropdownOpen = ref(false);


// =======================================================
// 3. DOM REFS
// =======================================================
const fileInput = ref(null);
const editor = ref(null);
const deptDropdownRef = ref(null);
const durationDropdownRef = ref(null);


// =======================================================
// 4. COMPUTED PROPERTIES
// =======================================================
/**
 * Checks if all required fields in Step 1 are filled.
 */
const isStep1Valid = computed(() => {
    // Check required text fields
    const isJobDetailsValid = jobForm.title.trim() !== '' && jobForm.description.trim() !== '';

    // Check application period: either a predefined option or a custom days > 0
    let isPeriodValid = false;
    if (selectedDateOption.value === 'custom') {
        isPeriodValid = customDays.value > 0;
    } else {
        isPeriodValid = DATE_OPTIONS.some(option => option.value === selectedDateOption.value && option.value !== 'custom');
    }

    return isJobDetailsValid && isPeriodValid;
});

/**
 * Check if there's any content to preview
 */
const hasPreviewContent = computed(() => {
    return jobForm.title.trim() !== '' ||
        jobForm.description.trim() !== '' ||
        uploadedImages.value.length > 0;
});


// =======================================================
// 5. WATCHERS
// =======================================================
/**
 * Syncs selectedDateOption to jobForm.applicationDays.
 */
watch(selectedDateOption, (newOption) => {
    if (newOption === 'custom') {
        // Use customDays value, ensure it's at least 1
        jobForm.applicationDays = customDays.value > 0 ? customDays.value : 1;
    } else {
        // Use predefined days
        const option = DATE_OPTIONS.find(o => o.value === newOption);
        jobForm.applicationDays = option ? option.days : 15;
    }
});

/**
 * Syncs customDays input to jobForm.applicationDays when 'custom' is selected.
 */
watch(customDays, (newDays) => {
    if (selectedDateOption.value === 'custom') {
        // Ensure newDays is a positive number, otherwise default to 1
        jobForm.applicationDays = newDays > 0 ? newDays : 1;
    }
});

/**
 * Reset image index and scroll to top when switching preview mode.
 */
watch(previewMode, () => {
    currentImageIndex.value = 0;
    window.scrollTo({ top: 0, behavior: 'smooth' });
});


// =======================================================
// 6. CORE METHODS & HANDLERS
// =======================================================
// A. Navigation & Core Actions
const goBack = () => {
    router.back();
};

const nextStep = () => {
    if (currentStep.value === 1 && !isStep1Valid.value) {
        alert('กรุณากรอกข้อมูลที่จำเป็นให้ครบถ้วน');
        return;
    }
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
    if (confirm('คุณต้องการยกเลิกการสร้างประกาศงานใช่หรือไม่? การดำเนินการทั้งหมดจะไม่ถูกบันทึก')) {
        router.back();
    }
};

const publishJob = () => {
    // TODO: Replace with actual API call
    console.log('Publishing job data:', {
        ...jobForm,
        images: uploadedImages.value.map(img => img.file.name),
        applicationPeriod: getDateRangePreview()
    });
    alert('ประกาศงานสำเร็จ!');
    router.push({ name: 'CompanyJobs' }); // Redirect
};

// B. Image Handling
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
                preview: e.target.result
            });
        };
        reader.readAsDataURL(file);
    });
    event.target.value = ''; // Clear file input
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

// C. Dropdown & UI Handlers
const selectDepartment = (dept) => {
    jobForm.department = dept;
    isDeptDropdownOpen.value = false;
};

const selectDurationOption = (option) => {
    selectedDateOption.value = option.value;
    isDurationDropdownOpen.value = false;
};

const handleClickOutside = (event) => {
    if (deptDropdownRef.value && !deptDropdownRef.value.contains(event.target)) {
        isDeptDropdownOpen.value = false;
    }
    if (durationDropdownRef.value && !durationDropdownRef.value.contains(event.target)) {
        isDurationDropdownOpen.value = false;
    }
};

// D. Rich Text Editor Logic
const execCommand = (command) => {
    document.execCommand(command, false, null);
    editor.value.focus();
    updateDescription();
};

const updateDescription = () => {
    jobForm.description = editor.value.innerHTML;
};

// E. Timeline Helper
const getConnectorClass = (index) => {
    const currentStage = jobStages.value[index];
    const nextStage = jobStages.value[index + 1];

    if (currentStage.status === "passed" && (nextStage.status === "passed" || nextStage.status === "current")) {
        return "passed";
    }

    return "pending";
};


// =======================================================
// 7. UTILITY FUNCTIONS
// =======================================================
/**
 * Formats a Date object into "DD Mon YYYY" (Thai Buddhist Year).
 */
const formatDate = (date) => {
    const d = new Date(date);
    const day = d.getDate();
    const months = [
        "ม.ค.", "ก.พ.", "มี.ค.", "เม.ย.", "พ.ค.", "มิ.ย.",
        "ก.ค.", "ส.ค.", "ก.ย.", "ต.ค.", "พ.ย.", "ธ.ค."
    ];
    const month = months[d.getMonth()];
    const year = d.getFullYear() + 543; // Convert to Thai Buddhist Year
    return `${day} ${month} ${year}`;
};

/**
 * Calculates and formats the application period display string for preview.
 */
const getDateRangePreview = () => {
    const days = jobForm.applicationDays;
    if (days > 0) {
        return `รออนุมัติจากอาจารย์ผู้ประสานงานที่เกี่ยวข้อง (ระยะเวลา ${days} วัน)`;
    }
    return 'ไม่ระบุระยะเวลา';
};

const truncateHTML = (html, maxLength) => {
    if (!html) return '';
    const text = html.replace(/<[^>]*>/g, ''); // ลบ HTML tags เพื่อนับความยาว
    if (text.length <= maxLength) return html;

    // ตัดข้อความและเติม ...
    const div = document.createElement('div');
    div.innerHTML = html;
    let truncated = div.textContent || div.innerText || '';
    truncated = truncated.substring(0, maxLength);
    return truncated;
};


// =======================================================
// 8. LIFECYCLE HOOKS
// =======================================================
onMounted(() => {
    document.addEventListener("click", handleClickOutside);
    if (editor.value) {
        jobForm.description = editor.value.innerHTML;
    }
});

onBeforeUnmount(() => {
    document.removeEventListener("click", handleClickOutside);
});

</script>

<style scoped>
/* =================================== */
/* GLOBAL STYLES & FONT (THIS IS YOUR ORIGINAL STYLE) */
/* =================================== */
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

/* Common Classes */
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
/* PROGRESS STEPS */
/* =================================== */
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

/* =================================== */
/* CARD & HELPERS */
/* =================================== */
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

.image-preview-item.active {
    border-color: #037266;
    box-shadow: 0 0 0 3px rgba(3, 114, 102, 0.2);
    transform: scale(1.05);
}

.image-preview-item {
    transition: all 0.2s ease;
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

/* =================================== */
/* CUSTOM DROPDOWN (Reusable) */
/* =================================== */
.custom-dropdown {
    position: relative;
    width: 100%;
    /* Changed from 250px for better layout control */
    min-width: 250px;
    cursor: pointer;
    user-select: none;
}

.dropdown-selected {
    display: flex;
    align-items: center;
    justify-content: space-between;
    padding: 14px 16px;
    /* Adjusted padding to match form-input */
    background: #fff;
    border: 2px solid #e5e5e5;
    /* Adjusted border to match form-input */
    border-radius: 10px;
    font-size: 15px;
    /* Adjusted font size */
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

/* Transition */
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
/* PERIOD DROPDOWN & CUSTOM INPUT */
/* =================================== */
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
    flex: 1;
}

.custom-input-wrapper {
    display: flex;
    align-items: center;
    gap: 10px;
    flex-shrink: 0;
}

.custom-days-input {
    width: 140px;
    padding: 14px 16px;
    border: 2px solid #e5e5e5;
    border-radius: 10px;
    font-size: 15px;
    color: #333;
    background: #fff;
    transition: all 0.2s ease;
    font-family: "Kanit", sans-serif;
    font-weight: 500;
}

.custom-days-input:focus {
    outline: none;
    border-color: #037266;
    background: #f0f8f7;
}

.days-suffix {
    font-size: 15px;
    font-weight: 500;
    color: #666;
    white-space: nowrap;
}

/* Tooltip Styles (Kept as is - good) */
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

.tooltip-content {
    position: absolute;
    top: 100%;
    right: 0px;
    background: #fffbf0;
    border: 2px solid #f4d669;
    border-radius: 12px;
    padding: 16px;
    margin-top: 8px;
    width: 380px;
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

.notice-highlight {
    color: #f59e0b !important;
    font-weight: 600 !important;
    margin-top: 10px;
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
/* CHECKBOX GROUP */
/* =================================== */
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

/* =================================== */
/* STEP 2: PREVIEW CARD */
/* =================================== */
.preview-card {
    background: #f9f9f9;
    border-radius: 16px;
    padding: 30px;
    box-shadow: 0 2px 12px rgba(3, 114, 102, 0.07);
    border: 1px solid #eef0ef;
    margin-bottom: 25px;
}

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

/* ========== CONFIRMATION CARD (STEP 2) ========== */
.confirmation-card {
    text-align: center;
    padding: 40px 30px;
    margin-top: 30px;
}

.confirmation-card h3 {
    display: flex;
    justify-content: center;
    width: 100%;
    font-size: 24px;
    margin-bottom: 20px;
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

.confirmation-notice {
    background: #fff9e6;
    border: 1px solid #ffd43b;
    border-radius: 12px;
    padding: 20px;
    display: flex;
    align-items: flex-start;
    gap: 15px;
    text-align: left;
}

.confirmation-notice svg {
    color: #f59e0b;
    flex-shrink: 0;
    margin-top: 2px;
}

.confirmation-notice p {
    font-size: 14px;
    color: #666;
    margin: 0;
    line-height: 1.6;
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

.btn-confirm:hover:not(:disabled) {
    background: linear-gradient(135deg, #059669 0%, #047857 100%);
}

/* =================================== */
/* ANIMATION */
/* =================================== */
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

/* =================================== */
/* PREVIEW TABS & CONTAINER */
/* =================================== */
.preview-tabs {
    display: flex;
    gap: 12px;
    margin-bottom: 20px;
    padding: 6px;
    background: #f7f9f8;
    border-radius: 12px;
    border: 1px solid #eef4f2;
}

.tab-btn {
    flex: 1;
    display: flex;
    align-items: center;
    justify-content: center;
    gap: 10px;
    padding: 12px 20px;
    border: none;
    border-radius: 10px;
    background: transparent;
    color: #666;
    font-size: 15px;
    font-weight: 600;
    cursor: pointer;
    transition: all 0.2s;
}

.tab-btn:hover {
    background: #fff;
    color: #037266;
}

.tab-btn.active {
    background: linear-gradient(135deg, #037266 0%, #026b5b 100%);
    color: #fff;
    box-shadow: 0 4px 12px rgba(3, 114, 102, 0.15);
}

.tab-btn svg {
    flex-shrink: 0;
}

.preview-container {
    animation: fadeIn 0.3s ease;
}

.preview-notice {
    display: flex;
    align-items: center;
    gap: 12px;
    padding: 14px 18px;
    background: #fffbf0;
    border: 2px solid #ffd43b;
    border-radius: 12px;
    margin-bottom: 20px;
    font-size: 14px;
    color: #666;
    font-weight: 500;
}

.preview-notice svg {
    color: #f59e0b;
    flex-shrink: 0;
}

/* =================================== */
/* JOB CARD PREVIEW (คัดลอกจาก JobsSections) */
/* =================================== */
.job-card-preview {
    background: #fff;
    border-radius: 16px;
    box-shadow: 0 2px 12px rgba(3, 114, 102, 0.07);
    border: 1px solid #eef0ef;
    padding: 25px 30px 22px 30px;
    display: flex;
    flex-direction: column;
    gap: 16px;
}

.job-card-preview .job-header {
    display: flex;
    align-items: flex-start;
    gap: 20px;
}

.job-card-preview .company-logo {
    width: 58px;
    height: 58px;
    border-radius: 50%;
    object-fit: cover;
    background: #f0f0f0;
    border: 1.5px solid #e1e1e1;
    flex-shrink: 0;
}

.job-card-preview .job-header-info {
    display: flex;
    flex-direction: column;
    gap: 6px;
    flex: 1;
}

.job-card-preview .company-name {
    font-size: 17px;
    font-weight: 700;
    color: #333;
    text-transform: uppercase;
    letter-spacing: 0.03em;
}

.job-card-preview .department-label {
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

.job-card-preview .job-meta {
    font-size: 13px;
    color: #6b6b6b;
    font-weight: 500;
    display: flex;
    flex-direction: column;
    gap: 3px;
    margin-top: 2px;
}

.job-card-preview .job-meta-row {
    display: flex;
    align-items: center;
    gap: 5px;
}

.job-card-preview .meta-icon {
    display: flex;
    align-items: center;
    justify-content: center;
    margin-right: 3px;
}

.job-card-preview .job-content {
    display: flex;
    gap: 28px;
    align-items: flex-start;
    margin-left: 75px;
}

.job-card-preview .job-images {
    min-width: 220px;
    max-width: 220px;
    flex-shrink: 0;
}

.job-card-preview .job-img {
    width: 220px;
    height: 160px;
    object-fit: cover;
    border-radius: 8px;
}

.job-card-preview .image-grid {
    display: grid;
    gap: 6px;
}

.job-card-preview .image-cell {
    position: relative;
}

.job-card-preview .image-cell img {
    width: 100%;
    height: 100%;
    border-radius: 8px;
    object-fit: cover;
}

.job-card-preview .image-overlay {
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

.job-card-preview .image-grid.images-2 {
    grid-template-columns: repeat(2, 1fr);
    grid-auto-rows: 160px;
    width: 220px;
}

.job-card-preview .image-grid.images-3 {
    grid-template-areas: "a b" "c c";
    grid-template-columns: repeat(2, 1fr);
    grid-template-rows: 80px 80px;
    width: 220px;
}

.job-card-preview .image-grid.images-3 .image-cell:nth-child(1) {
    grid-area: a;
}

.job-card-preview .image-grid.images-3 .image-cell:nth-child(2) {
    grid-area: b;
}

.job-card-preview .image-grid.images-3 .image-cell:nth-child(3) {
    grid-area: c;
}

.job-card-preview .image-grid.images-4 {
    grid-template-columns: repeat(2, 1fr);
    grid-template-rows: repeat(2, 80px);
    width: 220px;
}

.job-card-preview .job-details {
    flex: 1;
    display: flex;
    flex-direction: column;
    gap: 8px;
    min-width: 0;
}

.job-card-preview .job-title {
    font-size: 22px;
    font-weight: 700;
    color: #222;
    margin: 0;
}

.job-card-preview .job-details-title {
    font-size: 15px;
    font-weight: 600;
    color: #777;
    margin: 0;
}

.job-card-preview .job-description {
    font-size: 14px;
    color: #444;
    line-height: 1.6;
    font-weight: 400;
    white-space: pre-line;
    margin: 0;
}

.job-card-preview .show-more {
    color: #009688;
    font-weight: 600;
    font-size: 14px;
    margin-left: 4px;
}

/* =================================== */
/* SIDE PANEL PREVIEW */
/* =================================== */
.side-panel-preview {
    background: #fff;
    border-radius: 16px;
    box-shadow: 0 2px 12px rgba(3, 114, 102, 0.07);
    border: 1px solid #eef0ef;
    padding: 40px 30px;
    max-width: 900px;
    margin: 0 auto;
}

.side-panel-preview .side-panel-inner {
    padding: 0;
    display: flex;
    flex-direction: column;
    gap: 18px;
}

.side-panel-preview .side-header-top {
    display: flex;
    align-items: flex-start;
    gap: 26px;
    margin-bottom: 4px;
}

.side-panel-preview .side-logo {
    width: 120px;
    height: 120px;
    border-radius: 50%;
    object-fit: cover;
    background: #f3f3f4;
    border: 2.2px solid #e1e1e1;
    flex-shrink: 0;
    box-shadow: 0 2px 13px rgba(8, 12, 22, 0.11);
}

.side-panel-preview .side-title-block {
    display: flex;
    flex-direction: column;
    gap: 7px;
    flex: 1;
    min-width: 0;
}

.side-panel-preview .side-title {
    font-size: 24px;
    font-weight: 700;
    color: #222;
    margin-bottom: 2px;
    margin-top: 0;
    word-break: break-word;
}

.side-panel-preview .side-company-row {
    display: flex;
    align-items: center;
    gap: 10px;
    font-size: 14px;
    color: #666;
    margin-bottom: 0;
}

.side-panel-preview .side-company-name {
    font-weight: 600;
    font-size: 15px;
    color: #037266;
    text-transform: uppercase;
    letter-spacing: 0.04em;
}

.side-panel-preview .side-company-sep {
    width: 5px;
    height: 5px;
    border-radius: 50%;
    background: #d7d7d7;
    display: inline-block;
}

.side-panel-preview .side-company-all-jobs {
    color: #888;
    font-size: 14px;
}

.side-panel-preview .side-department-label {
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

.side-panel-preview .side-meta-row {
    display: flex;
    align-items: center;
    gap: 0px;
    margin: 7px 0 0 0;
    font-size: 14px;
    color: #667;
    flex-wrap: wrap;
    flex-direction: row;
}

.side-panel-preview .side-meta-row+.side-meta-row {
    margin-top: 0;
}

.side-panel-preview .side-meta-item {
    display: flex;
    align-items: center;
    gap: 7px;
    font-size: 15px;
    color: #3b5454;
    font-weight: 500;
    margin-bottom: 0.5px;
}

.side-panel-preview .side-meta-item svg {
    flex-shrink: 0;
    margin-bottom: -2px;
}

.side-panel-preview .side-images-preview {
    margin: 25px 0 11px 0;
    display: flex;
    flex-direction: row;
    gap: 11px;
}

.side-panel-preview .side-image-layout {
    display: flex;
    flex-direction: row;
    gap: 16px;
}

.side-panel-preview .side-image-large {
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

.side-panel-preview .side-image-large img {
    width: 95%;
    height: 95%;
    object-fit: cover;
    border-radius: 10px;
    background: #fff;
    box-shadow: 0 1px 8px rgba(0, 0, 0, 0.07);
}

.side-panel-preview .side-image-thumbs-wrapper {
    max-height: 420px;
    overflow-y: auto;
    overflow-x: hidden;
    flex-shrink: 0;
}

.side-panel-preview .side-image-thumbs-wrapper::-webkit-scrollbar {
    width: 6px;
}

.side-panel-preview .side-image-thumbs-wrapper::-webkit-scrollbar-track {
    background: #f1f1f1;
    border-radius: 10px;
}

.side-panel-preview .side-image-thumbs-wrapper::-webkit-scrollbar-thumb {
    background: #c1c1c1;
    border-radius: 10px;
}

.side-panel-preview .side-image-thumbs {
    display: flex;
    flex-direction: column;
    gap: 12px;
    align-items: flex-start;
    justify-content: flex-start;
    min-width: 96px;
}

.side-panel-preview .side-image-thumb {
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

.side-panel-preview .side-image-thumb img {
    width: 100%;
    height: 100%;
    object-fit: cover;
    border-radius: 6px;
}

.side-panel-preview .side-image-thumb.active {
    border-color: #037266;
    box-shadow: 0 3px 12px rgba(3, 114, 102, 0.18);
    transform: scale(1.03);
}

.side-panel-preview .side-image-thumb:hover {
    border-color: #037266;
    box-shadow: 0 3px 12px rgba(3, 114, 102, 0.14);
}

.side-panel-preview .side-detail-content {
    margin-top: 18px;
    display: flex;
    flex-direction: column;
    gap: 15px;
}

.side-panel-preview .side-detail-title {
    font-size: 18px;
    font-weight: 700;
    color: #222;
    margin-bottom: 3px;
}

.side-panel-preview .side-description {
    font-size: 15px;
    line-height: 1.7;
    color: #444;
    background: none;
    padding: 0;
    border-radius: 0;
    white-space: pre-line;
}

/* =================================== */
/* ANIMATIONS & TRANSITIONS */
/* =================================== */
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

@keyframes slideIn {
    from {
        opacity: 0;
        transform: translateX(-10px);
    }

    to {
        opacity: 1;
        transform: translateX(0);
    }
}

.preview-container {
    animation: fadeIn 0.3s ease;
}

.job-card-preview,
.side-panel-preview {
    animation: slideIn 0.4s ease;
}

/* Smooth transitions for interactive elements */
.side-panel-preview .side-image-thumb,
.job-card-preview .image-cell {
    transition: all 0.2s ease;
}

.side-panel-preview .side-image-thumb:hover,
.job-card-preview .image-cell:hover {
    transform: scale(1.02);
}


/* ========== APPLICATION SUMMARY (STEP 3) ========== */
.application-summary-card {
    background: #f9f9f9;
    border-radius: 16px;
    padding: 40px;
    box-shadow: 0 2px 12px rgba(3, 114, 102, 0.07);
    border: 1px solid #eef0ef;
    margin-bottom: 25px;
}

.summary-header {
    text-align: center;
    margin-bottom: 40px;
    padding-bottom: 30px;
    border-bottom: 2px solid #e0e0e0;
}

.summary-icon {
    width: 80px;
    height: 80px;
    background: linear-gradient(135deg, #037266 0%, #026b5b 100%);
    border-radius: 50%;
    display: flex;
    align-items: center;
    justify-content: center;
    margin: 0 auto 20px;
    color: #fff;
    box-shadow: 0 4px 16px rgba(3, 114, 102, 0.2);
}

.summary-header h2 {
    font-size: 26px;
    font-weight: 700;
    color: #222;
    margin: 0 0 10px 0;
}

.summary-subtitle {
    font-size: 15px;
    color: #666;
    margin: 0;
}

.summary-section {
    background: #fff;
    border-radius: 12px;
    padding: 25px;
    margin-bottom: 20px;
    box-shadow: 0 1px 4px rgba(0, 0, 0, 0.04);
}

.section-header-summary {
    display: flex;
    align-items: center;
    gap: 10px;
    margin-bottom: 20px;
    padding-bottom: 15px;
    border-bottom: 2px solid #f0f0f0;
}

.section-header-summary svg {
    color: #037266;
    flex-shrink: 0;
}

.section-header-summary h4 {
    font-size: 18px;
    font-weight: 600;
    color: #222;
    margin: 0;
}

.summary-content {
    padding: 10px 0;
}

.job-position-display {
    display: flex;
    align-items: center;
    gap: 20px;
}

.job-logo-small {
    width: 60px;
    height: 60px;
    border-radius: 50%;
    object-fit: cover;
    border: 2px solid #e1e1e1;
    flex-shrink: 0;
}

.job-details-summary {
    flex: 1;
}

.job-details-summary h5 {
    font-size: 18px;
    font-weight: 700;
    color: #222;
    margin: 0 0 6px 0;
}

.company-detail {
    font-size: 14px;
    color: #666;
    margin: 0 0 8px 0;
}

.department-badge {
    display: inline-block;
    background: #e0f2f0;
    color: #037266;
    padding: 4px 12px;
    border-radius: 12px;
    font-size: 12px;
    font-weight: 600;
}

.applicant-info-grid {
    display: flex;
    flex-direction: column;
    gap: 15px;
}

.info-row-summary {
    display: flex;
    flex-direction: column;
    gap: 6px;
    padding: 12px 0;
    border-bottom: 1px solid #f5f5f5;
}

.info-row-summary:last-child {
    border-bottom: none;
}

.info-label-summary {
    font-size: 13px;
    color: #999;
    font-weight: 500;
    text-transform: uppercase;
    letter-spacing: 0.5px;
}

.info-value-summary {
    font-size: 15px;
    color: #333;
    font-weight: 600;
    display: flex;
    align-items: center;
    gap: 8px;
}

.info-value-summary svg {
    color: #999;
    flex-shrink: 0;
}

.documents-list {
    display: flex;
    flex-direction: column;
    gap: 12px;
}

.document-item {
    display: flex;
    align-items: center;
    gap: 15px;
    padding: 15px;
    background: #f9f9f9;
    border: 1px solid #e5e5e5;
    border-radius: 10px;
}

.document-icon {
    width: 42px;
    height: 42px;
    border-radius: 10px;
    display: flex;
    align-items: center;
    justify-content: center;
    flex-shrink: 0;
}

.resume-icon {
    background: linear-gradient(135deg, #037266 0%, #026b5b 100%);
    color: #fff;
}

.document-info {
    flex: 1;
    display: flex;
    flex-direction: column;
    gap: 4px;
}

.document-type {
    font-size: 13px;
    color: #999;
    font-weight: 500;
}

.document-name {
    font-size: 15px;
    color: #333;
    font-weight: 600;
}

.application-date {
    display: flex;
    align-items: center;
    justify-content: center;
    gap: 8px;
    padding: 15px;
    background: #fff;
    border-radius: 10px;
    font-size: 14px;
    color: #666;
    font-weight: 500;
    margin-top: 10px;
    box-shadow: 0 1px 4px rgba(0, 0, 0, 0.04);
}

.application-date svg {
    color: #037266;
    flex-shrink: 0;
}

/* ========== STATUS TIMELINE CARD ========== */
.status-timeline-card {
    background: #fff;
    border-radius: 16px;
    padding: 30px;
    box-shadow: 0 2px 12px rgba(3, 114, 102, 0.07);
    border: 1px solid #eef0ef;
    margin-top: 25px;
}

.timeline-header {
    margin-bottom: 30px;
    padding-bottom: 20px;
    border-bottom: 2px solid #f0f0f0;
}

.timeline-header h3 {
    font-size: 20px;
    font-weight: 700;
    color: #222;
    margin: 0;
}

/* Timeline Container */
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

/* Timeline Dot */
.timeline-dot {
    width: 60px;
    height: 60px;
    border-radius: 50%;
    background: #f8f8f8;
    border: 3px solid #e8e8e8;
    display: flex;
    align-items: center;
    justify-content: center;
    margin-bottom: 18px;
    transition: all 0.3s ease;
    flex-shrink: 0;
    position: relative;
    z-index: 3;
    box-shadow: 0 2px 8px rgba(0, 0, 0, 0.08);
}

.timeline-dot.passed {
    background: linear-gradient(135deg, #10b981 0%, #059669 100%);
    border-color: #10b981;
    color: #fff;
    box-shadow: 0 4px 16px rgba(16, 185, 129, 0.3);
}

.timeline-dot.current {
    background: linear-gradient(135deg, #fbbf24 0%, #f59e0b 100%);
    border-color: #f59e0b;
    color: #fff;
    box-shadow: 0 0 0 6px rgba(245, 158, 11, 0.15), 0 4px 16px rgba(245, 158, 11, 0.3);
    animation: pulseJobStatus 2s infinite;
}

.timeline-dot.pending {
    background: #f8f8f8;
    border-color: #e8e8e8;
    color: #ccc;
    box-shadow: 0 2px 6px rgba(0, 0, 0, 0.05);
}

@keyframes pulseJobStatus {

    0%,
    100% {
        box-shadow: 0 0 0 6px rgba(245, 158, 11, 0.15), 0 4px 16px rgba(245, 158, 11, 0.3);
        transform: scale(1);
    }

    50% {
        box-shadow: 0 0 0 10px rgba(245, 158, 11, 0.1), 0 4px 16px rgba(245, 158, 11, 0.3);
        transform: scale(1.02);
    }
}

/* Timeline Content */
.timeline-content {
    display: flex;
    flex-direction: column;
    align-items: center;
    gap: 8px;
    text-align: center;
}

.timeline-label {
    font-size: 15px;
    color: #222;
    font-weight: 700;
    letter-spacing: 0.3px;
}

.timeline-description {
    font-size: 13px;
    color: #666;
    line-height: 1.5;
    max-width: 220px;
    margin: 0;
}

/* Timeline Connector */
.timeline-connector {
    position: absolute;
    top: 29px;
    left: calc(50% + 32px);
    right: calc(-50% + 32px);
    height: 4px;
    background: #e8e8e8;
    z-index: 1;
    transition: all 0.3s ease;
}

.timeline-connector.passed {
    background: linear-gradient(to right, #10b981 0%, #10b981 100%);
    box-shadow: 0 1px 6px rgba(16, 185, 129, 0.2);
}

.timeline-item:last-child .timeline-connector {
    display: none;
}

/* Timeline Info Box */
.timeline-info {
    margin-top: 25px;
}

.timeline-info .info-box {
    background: #f0f9ff;
    border: 1px solid #bae6fd;
    border-radius: 12px;
    padding: 18px 20px;
    display: flex;
    align-items: flex-start;
    gap: 12px;
}

.timeline-info .info-box svg {
    color: #0369a1;
    flex-shrink: 0;
    margin-top: 2px;
}

.timeline-info .info-box p {
    font-size: 14px;
    color: #0c4a6e;
    margin: 0;
    line-height: 1.6;
    font-weight: 500;
}

/* =================================== */
/* RESPONSIVE */
/* =================================== */
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

    .confirmation-notice {
        flex-direction: column;
        text-align: center;
    }

    .period-dropdown-wrapper {
        flex-wrap: wrap;
    }

    .period-select-container {
        width: 100%;
        min-width: unset;
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

    .status-timeline-card {
        padding: 20px;
    }

    .timeline-header {
        margin-bottom: 25px;
        padding-bottom: 15px;
    }

    .timeline-container {
        padding: 0 5px;
    }

    .timeline-wrapper {
        flex-direction: column;
        gap: 30px;
    }

    .timeline-item {
        width: 100%;
    }

    .timeline-dot {
        width: 55px;
        height: 55px;
    }

    .timeline-content {
        align-items: flex-start;
        text-align: left;
        margin-top: 5px;
    }

    .timeline-description {
        max-width: 100%;
    }

    .timeline-connector {
        top: 55px;
        left: 27px;
        right: auto;
        width: 4px;
        height: calc(100% - 25px);
    }

    .timeline-item:last-child .timeline-connector {
        display: none;
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

@media (max-width: 768px) {
    /* ... existing styles ... */

    /* Preview Tabs */
    .preview-tabs {
        flex-direction: column;
        gap: 8px;
    }

    .tab-btn {
        width: 100%;
        padding: 10px 16px;
    }

    /* Job Card Preview */
    .job-card-preview .job-content {
        flex-direction: column;
        gap: 20px;
        margin-left: 0;
    }

    .job-card-preview .job-images {
        max-width: 100%;
        min-width: 100%;
    }

    .job-card-preview .job-img {
        width: 100%;
    }

    .job-card-preview .image-grid.images-2,
    .job-card-preview .image-grid.images-3,
    .job-card-preview .image-grid.images-4 {
        width: 100%;
    }

    /* Side Panel Preview */
    .side-panel-preview {
        padding: 25px 20px;
    }

    .side-panel-preview .side-image-layout {
        flex-direction: column;
        gap: 16px;
        align-items: center;
    }

    .side-panel-preview .side-image-large {
        width: 100%;
        max-width: 100%;
        height: 300px;
    }

    .side-panel-preview .side-image-thumbs-wrapper {
        height: auto;
        max-height: none;
        overflow-y: visible;
        overflow-x: auto;
        width: 100%;
        max-width: 100%;
    }

    .side-panel-preview .side-image-thumbs {
        flex-direction: row;
        gap: 10px;
        width: max-content;
    }

    .side-panel-preview .side-image-thumb {
        width: 80px;
        height: 80px;
    }
}

@media (max-width: 480px) {
    /* ... existing styles ... */

    /* Job Card Preview */
    .job-card-preview {
        padding: 18px 16px;
    }

    .job-card-preview .company-logo {
        width: 46px;
        height: 46px;
    }

    .job-card-preview .job-header {
        gap: 14px;
    }

    .job-card-preview .job-title {
        font-size: 18px;
    }

    /* Side Panel Preview */
    .side-panel-preview {
        padding: 20px 15px;
    }

    .side-panel-preview .side-logo {
        width: 70px;
        height: 70px;
    }

    .side-panel-preview .side-title {
        font-size: 20px;
    }

    .side-panel-preview .side-image-large {
        height: 240px;
    }

    .side-panel-preview .side-image-thumb {
        width: 65px;
        height: 65px;
    }
}

/* =================================== */
/* EMPTY STATE */
/* =================================== */
.empty-preview-state {
    background: #fff;
    border-radius: 16px;
    border: 2px dashed #e0e0e0;
    padding: 60px 40px;
    text-align: center;
    display: flex;
    flex-direction: column;
    align-items: center;
    gap: 20px;
    margin-bottom: 25px;
}

.empty-preview-state svg {
    color: #ccc;
}

.empty-preview-state h3 {
    font-size: 20px;
    font-weight: 600;
    color: #555;
    margin: 0;
}

.empty-preview-state p {
    font-size: 15px;
    color: #888;
    margin: 0;
    max-width: 400px;
    line-height: 1.6;
}

.empty-preview-state button {
    margin-top: 10px;
    display: inline-flex;
    align-items: center;
    gap: 8px;
}

/* =================================== */
/* IMAGE LOADING STATE */
/* =================================== */
.job-card-preview .job-img,
.job-card-preview .image-cell img,
.side-panel-preview .side-image-large img,
.side-panel-preview .side-image-thumb img {
    background: linear-gradient(90deg, #f0f0f0 25%, #e0e0e0 50%, #f0f0f0 75%);
    background-size: 200% 100%;
    animation: shimmer 1.5s infinite;
}

.job-card-preview .job-img[src],
.job-card-preview .image-cell img[src],
.side-panel-preview .side-image-large img[src],
.side-panel-preview .side-image-thumb img[src] {
    background: none;
    animation: none;
}

@keyframes shimmer {
    0% {
        background-position: 200% 0;
    }

    100% {
        background-position: -200% 0;
    }
}

.preview-notice {
    position: relative;
    padding-right: 50px;
    /* Make room for dismiss button */
}

.notice-dismiss {
    position: absolute;
    right: 14px;
    top: 50%;
    transform: translateY(-50%);
    width: 28px;
    height: 28px;
    border: none;
    border-radius: 50%;
    background: transparent;
    color: #999;
    cursor: pointer;
    display: flex;
    align-items: center;
    justify-content: center;
    transition: all 0.2s;
}

.notice-dismiss:hover {
    background: rgba(0, 0, 0, 0.05);
    color: #666;
}
</style>