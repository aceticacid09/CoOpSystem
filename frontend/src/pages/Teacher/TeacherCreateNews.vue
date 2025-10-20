<!-- frontend/src/pages/Teacher/TeacherCreateNews.vue -->
<template>
    <DashboardLayout role="teacher">
        <div class="create-news-page">
            <div class="page-header">
                <button class="btn-back" @click="goBack">
                    <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none"
                        stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                        <line x1="19" y1="12" x2="5" y2="12"></line>
                        <polyline points="12 19 5 12 12 5"></polyline>
                    </svg>
                    ย้อนกลับ
                </button>
                <h1>สร้างข่าวสารและกิจกรรม</h1>
                <p class="page-subtitle">สร้างข่าวประชาสัมพันธ์และกิจกรรมเพื่อแจ้งให้นักศึกษาและสถานประกอบการทราบ</p>
            </div>

            <div class="progress-steps">
                <div class="step" :class="{ active: currentStep >= 1, completed: currentStep > 1 }">
                    <div class="step-number">1</div>
                    <div class="step-label">สร้างข่าวสาร</div>
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
                    <p class="helper-text">เขียนรายละเอียดเกี่ยวกับข่าวสารหรือกิจกรรมที่ต้องการประชาสัมพันธ์</p>

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
                    <p class="helper-text">เลือกเวลาที่ต้องการประกาศข่าวสาร</p>

                    <div class="form-group">
                        <div class="form-group">
                            <div class="publish-options">
                                <!-- ⭐ ตัวเลือกที่ 1: แบบร่าง (ใหม่) -->
                                <label class="radio-item">
                                    <input type="radio" v-model="publishType" value="draft" name="publishType" />
                                    <div class="radio-content">
                                        <div class="radio-box"></div>
                                        <div class="radio-info">
                                            <span class="radio-label">บันทึกเป็นแบบร่าง</span>
                                            <span class="radio-description">เก็บข่าวสารไว้โดยยังไม่เผยแพร่
                                                สามารถกลับมาแก้ไขและเผยแพร่ทีหลังได้</span>
                                        </div>
                                    </div>
                                </label>

                                <!-- ตัวเลือกที่ 2: ประกาศทันที -->
                                <label class="radio-item">
                                    <input type="radio" v-model="publishType" value="immediate" name="publishType" />
                                    <div class="radio-content">
                                        <div class="radio-box"></div>
                                        <div class="radio-info">
                                            <span class="radio-label">ประกาศทันที</span>
                                            <span
                                                class="radio-description">ข่าวสารจะถูกเผยแพร่ทันทีหลังจากกดยืนยัน</span>
                                        </div>
                                    </div>
                                </label>

                                <!-- ตัวเลือกที่ 3: กำหนดวันเวลา -->
                                <label class="radio-item">
                                    <input type="radio" v-model="publishType" value="scheduled" name="publishType" />
                                    <div class="radio-content">
                                        <div class="radio-box"></div>
                                        <div class="radio-info">
                                            <span class="radio-label">กำหนดวันเวลาประกาศ</span>
                                            <span
                                                class="radio-description">เลือกวันและเวลาที่ต้องการให้ประกาศข่าวสาร</span>
                                        </div>
                                    </div>
                                </label>
                            </div>

                            <!-- แสดงช่องกรอกวันที่และเวลาเมื่อเลือก "กำหนดวันเวลา" -->
                            <div v-if="publishType === 'scheduled'" class="schedule-inputs">
                                <div class="datetime-group">
                                    <div class="input-wrapper">
                                        <label>วันที่ประกาศ</label>
                                        <input type="date" v-model="newsForm.publishDate" class="form-input"
                                            :min="minDate" />
                                    </div>
                                    <div class="input-wrapper">
                                        <label>เวลาที่ประกาศ</label>
                                        <input type="time" v-model="newsForm.publishTime" class="form-input" />
                                    </div>
                                </div>
                            </div>
                        </div>
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
                    <div class="news-card-preview">
                        <div class="news-img" v-if="uploadedImages.length > 0">
                            <img v-if="uploadedImages.length === 1" :src="uploadedImages[0].preview" alt="news" />
                            <div v-else class="image-grid" :class="`images-${Math.min(uploadedImages.length, 4)}`">
                                <div v-for="(img, i) in uploadedImages.slice(0, 4)" :key="i" class="image-cell"
                                    @click="currentImageIndex = i" style="cursor: pointer;">
                                    <img :src="img.preview" alt="news" />
                                    <div v-if="i === 3 && uploadedImages.length > 4" class="image-overlay">
                                        +{{ uploadedImages.length - 4 }}
                                    </div>
                                </div>
                            </div>
                        </div>

                        <div class="news-content">
                            <h3 class="news-title">{{ newsForm.title || 'หัวข้อข่าว' }}</h3>
                            <p class="news-date">
                                <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" width="14" height="14"
                                    fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round"
                                    stroke-linejoin="round">
                                    <rect x="3" y="4" width="18" height="18" rx="2" ry="2" />
                                    <line x1="16" y1="2" x2="16" y2="6" />
                                    <line x1="8" y1="2" x2="8" y2="6" />
                                    <line x1="3" y1="10" x2="21" y2="10" />
                                </svg>
                                ประกาศเมื่อ {{ getPublishDatePreview() }}
                            </p>
                            <p class="news-desc" v-html="truncateHTML(newsForm.description, 160)"></p>
                            <span v-if="newsForm.description.length > 160" class="show-more">... ดูเพิ่มเติม</span>
                        </div>
                    </div>
                </div>

                <div v-if="previewMode === 'panel'" class="preview-container">
                    <div class="side-panel-preview">
                        <div class="side-panel-inner">
                            <div class="side-header">
                                <h2>{{ newsForm.title || 'หัวข้อข่าว' }}</h2>
                                <p class="news-date">
                                    <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" width="14" height="14"
                                        fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round"
                                        stroke-linejoin="round">
                                        <rect x="3" y="4" width="18" height="18" rx="2" ry="2" />
                                        <line x1="16" y1="2" x2="16" y2="6" />
                                        <line x1="8" y1="2" x2="8" y2="6" />
                                        <line x1="3" y1="10" x2="21" y2="10" />
                                    </svg>
                                    {{ getPublishDatePreview() }}
                                </p>
                            </div>

                            <p class="side-description" v-html="newsForm.description || 'ไม่มีรายละเอียด'"></p>

                            <div v-if="uploadedImages.length > 0" class="side-images">
                                <img v-for="(img, i) in uploadedImages" :key="i" :src="img.preview" alt="news" />
                            </div>
                        </div>
                    </div>
                </div>
                
                <div class="section-card confirmation-card">
                    <h3 style="text-align: center;">ยืนยันการประกาศข่าวสาร</h3>
                    <p class="confirmation-text">
                        คุณต้องการยืนยันการประกาศข่าวสาร
                        <strong>{{ newsForm.title }}</strong> ใช่หรือไม่?
                    </p>

                    <div class="confirmation-notice">
                        <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none"
                            stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                            <circle cx="12" cy="12" r="10"></circle>
                            <line x1="12" y1="8" x2="12" y2="12"></line>
                            <line x1="12" y1="16" x2="12.01" y2="16"></line>
                        </svg>
                        <p>
                            {{ publishType === 'draft'
                                ? 'ข่าวสารจะถูกบันทึกเป็นแบบร่าง และไม่มีการเผยแพร่'
                                : publishType === 'immediate'
                                    ? 'เมื่อกดยืนยันแล้ว ข่าวสารจะถูกเผยแพร่ทันที'
                                    : `ข่าวสารจะถูกเผยแพร่ในวันที่ ${formatScheduledDate()}`
                            }}
                        </p>
                    </div>
                </div>

                <div class="action-buttons">
                    <button class="btn-secondary" @click="prevStep">ย้อนกลับ</button>
                    <button class="btn-primary" @click="nextStep">ยืนยันและประกาศ</button>
                </div>
            </div>

            <div v-if="currentStep === 3" class="step-content">
                <div class="application-summary-card">
                    <div class="summary-header">
                        <div class="summary-icon">
                            <svg xmlns="http://www.w3.org/2000/svg" width="48" height="48" viewBox="0 0 24 24"
                                fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round"
                                stroke-linejoin="round">
                                <path d="M22 11.08V12a10 10 0 1 1-5.93-9.14"></path>
                                <polyline points="22 4 12 14.01 9 11.01"></polyline>
                            </svg>
                        </div>
                        <h2>ประกาศข่าวสารสำเร็จเรียบร้อยแล้ว</h2>
                        <p class="summary-subtitle">กรุณาตรวจสอบรายละเอียดข่าวสารของคุณ</p>
                    </div>

                    <div class="summary-section">
                        <div class="section-header-summary">
                            <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24"
                                fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round"
                                stroke-linejoin="round">
                                <path d="M14 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8z"></path>
                                <polyline points="14 2 14 8 20 8"></polyline>
                            </svg>
                            <h4>รายละเอียดข่าวสาร</h4>
                        </div>
                        <div class="summary-content">
                            <div class="applicant-info-grid">
                                <div class="info-row-summary">
                                    <span class="info-label-summary">หัวข้อข่าว</span>
                                    <span class="info-value-summary">{{ newsForm.title }}</span>
                                </div>
                                <div class="info-row-summary">
                                    <span class="info-label-summary">หมวดหมู่</span>
                                    <span class="info-value-summary">{{ newsForm.category || 'ไม่ระบุหมวดหมู่' }}</span>
                                </div>
                                <div class="info-row-summary">
                                    <span class="info-label-summary">วันที่ประกาศ</span>
                                    <span class="info-value-summary">{{ getPublishDatePreview() }}</span>
                                </div>
                            </div>
                        </div>
                    </div>
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
                        <span>สร้างข่าวสารเมื่อวันที่ {{ formatDate(new Date()) }}</span>
                    </div>
                </div>

                <div class="success-info-card">
                    <div class="info-box">
                        <svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none"
                            stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                            <circle cx="12" cy="12" r="10"></circle>
                            <line x1="12" y1="8" x2="12" y2="12"></line>
                            <line x1="12" y1="16" x2="12.01" y2="16"></line>
                        </svg>
                        <p>
                            {{ publishType === 'draft' 
                                ? 'ข่าวสารของคุณได้รับการบันทึกเป็นแบบร่างแล้ว คุณสามารถกลับมาแก้ไขและเผยแพร่ได้ในภายหลังผ่านเมนู "จัดการข่าวสาร"' 
                                : publishType === 'immediate' 
                                    ? 'ข่าวสารของคุณได้รับการเผยแพร่แล้ว นักศึกษาและสถานประกอบการสามารถเห็นข่าวสารนี้ได้ทันที' 
                                    : `ข่าวสารของคุณจะถูกเผยแพร่อัตโนมัติในวันที่ ${formatScheduledDate()}` 
                            }}
                        </p>
                    </div>
                </div>

                <div class="action-buttons">
                    <button class="btn-secondary"
                        @click="router.push({ name: 'TeacherNewsEvents' })">กลับสู่หน้าข่าวสาร</button>
                    <button class="btn-primary" @click="createAnotherNews">สร้างข่าวสารใหม่</button>
                </div>
            </div>
        </div>
    </DashboardLayout>
</template>

<script setup>
import { ref, reactive, computed, watch, onMounted, onBeforeUnmount } from 'vue';
import { useRouter, useRoute } from 'vue-router';
import DashboardLayout from '../../components/DashboardLayout.vue';

// =======================================================
// 1. CONSTANTS & STATIC DATA
// =======================================================
const categories = [
    "ประชาสัมพันธ์",
    "ข่าวด่วน",
    "กิจกรรม",
    "ประกาศผลการคัดเลือก"
];

// =======================================================
// 2. STATE MANAGEMENT (REFS & REACTIVES)
// =======================================================
const router = useRouter();
const route = useRoute();

// A. Navigation & Flow State
const currentStep = ref(1);
const previewMode = ref('card');

// B. Form Data (Reactive)
const newsForm = reactive({
    title: '',
    category: '',
    description: '',
    publishDate: '',
    publishTime: ''
});

// C. UI & Dropdown State
const uploadedImages = ref([]);
const currentImageIndex = ref(0);
const publishType = ref('draft');
const isCategoryDropdownOpen = ref(false);

// =======================================================
// 3. DOM REFS
// =======================================================
const fileInput = ref(null);
const editor = ref(null);
const categoryDropdownRef = ref(null);

// =======================================================
// 4. COMPUTED PROPERTIES
// =======================================================
const isStep1Valid = computed(() => {
    const hasBasicInfo = newsForm.title.trim() !== '' &&
        newsForm.category.trim() !== '' &&
        newsForm.description.trim() !== '';

    // ⭐ ถ้าเลือกแบบร่าง: ต้องมีข้อมูลพื้นฐานเท่านั้น (ไม่บังคับวันที่)
    if (publishType.value === 'draft') {
        return hasBasicInfo;
    }

    // ถ้าเลือกกำหนดวันเวลา: ต้องมีข้อมูลพื้นฐาน + วันที่และเวลา
    if (publishType.value === 'scheduled') {
        return hasBasicInfo && newsForm.publishDate && newsForm.publishTime;
    }

    // ถ้าเลือกประกาศทันที: ต้องมีข้อมูลพื้นฐานเท่านั้น
    return hasBasicInfo;
});

const minDate = computed(() => {
    const today = new Date();
    return today.toISOString().split('T')[0];
});

// =======================================================
// 5. WATCHERS
// =======================================================
watch(previewMode, () => {
    currentImageIndex.value = 0;
    window.scrollTo({ top: 0, behavior: 'smooth' });
});

watch(publishType, (newType) => {
    // ⭐ ล้างวันที่เมื่อไม่ได้เลือก "กำหนดวันเวลา"
    if (newType !== 'scheduled') {
        newsForm.publishDate = '';
        newsForm.publishTime = '';
    }
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
    if (confirm('คุณต้องการยกเลิกการสร้างข่าวสารใช่หรือไม่? การดำเนินการทั้งหมดจะไม่ถูกบันทึก')) {
        router.back();
    }
};

const createAnotherNews = () => {
    // Reset form
    Object.assign(newsForm, {
        title: '',
        category: '',
        description: '',
        publishDate: '',
        publishTime: ''
    });
    uploadedImages.value = [];
    currentImageIndex.value = 0;
    publishType.value = 'immediate';
    currentStep.value = 1;

    if (editor.value) {
        editor.value.innerHTML = '';
    }

    window.scrollTo({ top: 0, behavior: 'smooth' });
};

// ⭐ ฟังก์ชันโหลดข้อมูลแบบร่าง
const loadDraftNews = (draftId) => {
    // นำเข้า newsList จาก newsData
    import('../../data/newsData.js').then((module) => {
        const foundNews = module.newsList.find(n => n.id === draftId && n.status === 'draft');
        
        if (foundNews) {
            // โหลดข้อมูลเข้า form
            newsForm.title = foundNews.title;
            newsForm.category = foundNews.category;
            newsForm.description = foundNews.description;
            
            // ตั้งค่า publishType
            publishType.value = 'draft';
            
            // โหลดรูปภาพถ้ามี
            if (foundNews.images && foundNews.images.length > 0) {
                uploadedImages.value = foundNews.images.map(img => ({
                    preview: img,
                    isExisting: true
                }));
            }
            
            // โหลด description เข้า editor
            if (editor.value) {
                editor.value.innerHTML = foundNews.description;
            }
        }
    });
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

// C. Dropdown & UI Handlers
const selectCategory = (cat) => {
    newsForm.category = cat;
    isCategoryDropdownOpen.value = false;
};

const handleClickOutside = (event) => {
    if (categoryDropdownRef.value && !categoryDropdownRef.value.contains(event.target)) {
        isCategoryDropdownOpen.value = false;
    }
};

// D. Rich Text Editor Logic
const execCommand = (command) => {
    document.execCommand(command, false, null);
    editor.value.focus();
    updateDescription();
};

const updateDescription = () => {
    newsForm.description = editor.value.innerHTML;
};

// =======================================================
// 7. UTILITY FUNCTIONS
// =======================================================
const formatDate = (date) => {
    const d = new Date(date);
    const day = d.getDate();
    const months = [
        "ม.ค.", "ก.พ.", "มี.ค.", "เม.ย.", "พ.ค.", "มิ.ย.",
        "ก.ค.", "ส.ค.", "ก.ย.", "ต.ค.", "พ.ย.", "ธ.ค."
    ];
    const month = months[d.getMonth()];
    const year = d.getFullYear() + 543;
    return `${day} ${month} ${year}`;
};

const formatFullDate = (date) => {
    const d = new Date(date);
    const day = d.getDate();
    const monthNames = [
        "มกราคม", "กุมภาพันธ์", "มีนาคม", "เมษายน", "พฤษภาคม", "มิถุนายน",
        "กรกฎาคม", "สิงหาคม", "กันยายน", "ตุลาคม", "พฤศจิกายน", "ธันวาคม",
    ];
    const month = monthNames[d.getMonth()];
    const year = d.getFullYear() + 543;
    return `${day} ${month} ${year}`;
};

const getPublishDatePreview = () => {
    if (publishType.value === 'immediate') {
        return formatDate(new Date());
    } else if (newsForm.publishDate) {
        return formatDate(new Date(newsForm.publishDate));
    }
    return 'ยังไม่ได้กำหนดวันที่';
};

const formatScheduledDate = () => {
    if (newsForm.publishDate) {
        const dateStr = formatFullDate(new Date(newsForm.publishDate));
        const timeStr = newsForm.publishTime || '00:00';
        return `${dateStr} เวลา ${timeStr} น.`;
    }
    return 'ยังไม่ได้กำหนดวันที่';
};

const truncateHTML = (html, maxLength) => {
    if (!html) return '';
    const text = html.replace(/<[^>]*>/g, '');
    if (text.length <= maxLength) return html;

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
        newsForm.description = editor.value.innerHTML;
    }
    
    // ⭐ ตรวจสอบว่ามี draftId ใน query หรือไม่
    const draftId = route.query.draftId;
    if (draftId) {
        loadDraftNews(parseInt(draftId));
    }
});

onBeforeUnmount(() => {
    document.removeEventListener("click", handleClickOutside);
});

</script>

<style scoped>
@import url("https://fonts.googleapis.com/css2?family=Kanit:wght@400;500;600;700&display=swap");

* {
    font-family: "Kanit", sans-serif;
    box-sizing: border-box;
}

.create-news-page {
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

.upload-small-text {
    font-size: 13px;
    color: #666;
    margin-top: 6px;
    font-weight: 500;
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

.radio-item input:checked+.radio-content {
    border-color: #037266;
    background: #f0f8f7;
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

.radio-item input:checked+.radio-content .radio-box {
    border-color: #037266;
    background: #fff;
}

.radio-item input:checked+.radio-content .radio-box::after {
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

.radio-description {
    font-size: 13px;
    color: #777;
    line-height: 1.5;
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
/* STEP 2: PREVIEW TABS & CONTAINER */
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

/* =================================== */
/* NEWS CARD PREVIEW (from NewsSections) */
/* =================================== */
.news-card-preview {
    background: #fff;
    border-radius: 16px;
    box-shadow: 0 2px 12px rgba(3, 114, 102, 0.07);
    border: 1px solid #eef0ef;
    padding: 20px;
    display: flex;
    gap: 20px;
    flex-wrap: wrap;
}

.news-card-preview .news-img {
    min-width: 220px;
    max-width: 220px;
    flex-shrink: 0;
}

.news-card-preview .news-img img {
    width: 220px;
    height: 160px;
    object-fit: cover;
    border-radius: 8px;
}

.news-card-preview .image-grid {
    display: grid;
    gap: 6px;
}

.news-card-preview .image-cell {
    position: relative;
}

.news-card-preview .image-cell img {
    width: 100%;
    height: 100%;
    border-radius: 8px;
    object-fit: cover;
}

.news-card-preview .image-overlay {
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

.news-card-preview .image-grid.images-2 {
    grid-template-columns: repeat(2, 1fr);
    grid-auto-rows: 160px;
    width: 220px;
}

.news-card-preview .image-grid.images-3 {
    grid-template-areas: "a b" "c c";
    grid-template-columns: repeat(2, 1fr);
    grid-template-rows: 80px 80px;
    width: 220px;
}

.news-card-preview .image-grid.images-3 .image-cell:nth-child(1) {
    grid-area: a;
}

.news-card-preview .image-grid.images-3 .image-cell:nth-child(2) {
    grid-area: b;
}

.news-card-preview .image-grid.images-3 .image-cell:nth-child(3) {
    grid-area: c;
}

.news-card-preview .image-grid.images-4 {
    grid-template-columns: repeat(2, 1fr);
    grid-template-rows: repeat(2, 80px);
    width: 220px;
}

.news-card-preview .news-content {
    flex: 1;
    display: flex;
    flex-direction: column;
    gap: 8px;
    min-width: 0;
}

.news-card-preview .news-title {
    font-size: 18px;
    color: #000;
    font-weight: 600;
    margin: 0;
}

.news-card-preview .news-date {
    display: inline-flex;
    align-items: center;
    gap: 6px;
    font-size: 13px;
    color: #555;
    margin: 0;
}

.news-card-preview .news-date svg {
    flex-shrink: 0;
}

.news-card-preview .news-desc {
    font-size: 14px;
    color: #444;
    line-height: 1.6;
    margin: 0;
}

.news-card-preview .show-more {
    color: #009688;
    font-weight: 600;
    font-size: 14px;
    margin-left: 4px;
}

/* =================================== */
/* SIDE PANEL PREVIEW (from NewsSections) */
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

.side-panel-preview .side-header h2 {
    font-size: 24px;
    font-weight: 700;
    margin-bottom: 8px;
    color: #222;
}

.side-panel-preview .side-header .news-date {
    display: inline-flex;
    align-items: center;
    gap: 6px;
    font-size: 13px;
    color: #555;
    margin: 0;
}

.side-panel-preview .side-description {
    margin-top: 15px;
    font-size: 15px;
    line-height: 1.7;
    color: #444;
    white-space: pre-line;
}

.side-panel-preview .side-images {
    display: flex;
    flex-wrap: wrap;
    gap: 12px;
    margin-top: 20px;
}

.side-panel-preview .side-images img {
    width: calc(50% - 6px);
    height: 150px;
    object-fit: cover;
    border-radius: 10px;
    transition: transform 0.25s ease;
}

.side-panel-preview .side-images img:hover {
    transform: scale(1.05);
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
/* CONFIRMATION CARD */
/* =================================== */
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
/* STEP 3: APPLICATION SUMMARY */
/* =================================== */
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

/* =================================== */
/* SUCCESS INFO CARD */
/* =================================== */
.success-info-card {
    background: #fff;
    border-radius: 16px;
    padding: 30px;
    box-shadow: 0 2px 12px rgba(3, 114, 102, 0.07);
    border: 1px solid #eef0ef;
    margin-top: 25px;
}

.info-box {
    background: #f0f9ff;
    border: 1px solid #bae6fd;
    border-radius: 12px;
    padding: 18px 20px;
    display: flex;
    align-items: flex-start;
    gap: 12px;
}

.info-box svg {
    color: #0369a1;
    flex-shrink: 0;
    margin-top: 2px;
}

.info-box p {
    font-size: 14px;
    color: #0c4a6e;
    margin: 0;
    line-height: 1.6;
    font-weight: 500;
}

/* =================================== */
/* ANIMATIONS */
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
/* RESPONSIVE */
/* =================================== */
@media (max-width: 768px) {
    .create-news-page {
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
    .confirmation-card,
    .application-summary-card,
    .success-info-card {
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

    .datetime-group {
        grid-template-columns: 1fr;
    }

    .news-card-preview {
        flex-direction: column;
        padding: 16px;
    }

    .news-card-preview .news-img {
        max-width: 100%;
        min-width: 100%;
    }

    .news-card-preview .news-img img {
        width: 100%;
    }

    .news-card-preview .image-grid.images-2,
    .news-card-preview .image-grid.images-3,
    .news-card-preview .image-grid.images-4 {
        width: 100%;
    }

    .side-panel-preview {
        padding: 25px 20px;
    }

    .side-panel-preview .side-images img {
        width: 100%;
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

    .confirmation-card {
        padding: 30px 20px;
    }

    .confirmation-card h3 {
        font-size: 20px;
    }

    .preview-tabs {
        flex-direction: column;
        gap: 8px;
    }

    .tab-btn {
        width: 100%;
        padding: 10px 16px;
    }
}
</style>