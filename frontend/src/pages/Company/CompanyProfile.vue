<!-- frontend/src/pages/Company/CompanyProfile.vue -->
<template>
  <DashboardLayout role="company">
    <section class="profile-section container">
      <!-- Tabs -->
      <TabNavigation v-model="activeTab" :tabs="tabs" />

      <!-- ========== ข้อมูลสถานประกอบการ ========== -->
      <div v-if="activeTab === 'ข้อมูลสถานประกอบการ'" class="application-content">
        <!-- ข้อมูลบริษัท Section -->
        <div class="section-wrapper">
          <h3 class="section-title-outside">ข้อมูลสถานประกอบการ</h3>
          <div class="section-card">
            <div class="company-profile-container">
              <!-- Company Logo -->
              <div class="company-logo-section">
                <div class="company-logo">
                  <img v-if="companyData.logoUrl" :src="companyData.logoUrl" alt="Company Logo" />
                  <svg v-else xmlns="http://www.w3.org/2000/svg" width="80" height="80" viewBox="0 0 24 24" fill="none"
                    stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                    <path d="M3 9l9-7 9 7v11a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2z"></path>
                    <polyline points="9 22 9 12 15 12 15 22"></polyline>
                  </svg>
                </div>
              </div>

              <!-- Company Info -->
              <div class="company-info-section">
                <h4 class="company-name">{{ companyData.companyName }}</h4>
                
                <div class="info-items">
                  <p class="info-item">
                    <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none"
                      stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                      <path d="M4 4h16c1.1 0 2 .9 2 2v12c0 1.1-.9 2-2 2H4c-1.1 0-2-.9-2-2V6c0-1.1.9-2 2-2z">
                      </path>
                      <polyline points="22,6 12,13 2,6"></polyline>
                    </svg>
                    <span>{{ companyData.email }}</span>
                  </p>
                  <p class="info-item">
                    <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none"
                      stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                      <path
                        d="M22 16.92v3a2 2 0 0 1-2.18 2 19.79 19.79 0 0 1-8.63-3.07 19.5 19.5 0 0 1-6-6 19.79 19.79 0 0 1-3.07-8.67A2 2 0 0 1 4.11 2h3a2 2 0 0 1 2 1.72 12.84 12.84 0 0 0 .7 2.81 2 2 0 0 1-.45 2.11L8.09 9.91a16 16 0 0 0 6 6l1.27-1.27a2 2 0 0 1 2.11-.45 12.84 12.84 0 0 0 2.81.7A2 2 0 0 1 22 16.92z">
                      </path>
                    </svg>
                    <span>{{ companyData.phone }}</span>
                  </p>
                </div>

                <button class="btn-edit" @click="openEditCompanyPanel">แก้ไข</button>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- ========== SIDE PANEL: แก้ไขข้อมูลบริษัท ========== -->
      <div v-if="showEditCompanyPanel" class="side-panel-overlay" @click.self="closeEditCompanyPanel">
        <div class="side-panel">
          <button class="side-close" @click="closeEditCompanyPanel">&times;</button>
          <div class="side-panel-inner">
            <h3 class="side-panel-title">แก้ไขข้อมูลสถานประกอบการ</h3>

            <!-- Logo Upload Section -->
            <div class="form-group">
              <label>โลโก้บริษัท</label>
              <div class="logo-upload-section">
                <div class="current-logo">
                  <img v-if="editCompanyData.logoUrl" :src="editCompanyData.logoUrl" alt="Current Logo" />
                  <div v-else class="no-logo">
                    <svg xmlns="http://www.w3.org/2000/svg" width="40" height="40" viewBox="0 0 24 24" fill="none"
                      stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                      <path d="M3 9l9-7 9 7v11a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2z"></path>
                      <polyline points="9 22 9 12 15 12 15 22"></polyline>
                    </svg>
                  </div>
                </div>
                <div class="logo-actions">
                  <input type="file" ref="logoInput" accept="image/*" @change="handleLogoUpload" style="display: none;" />
                  <button type="button" class="btn-upload-logo" @click="triggerLogoInput">
                    <svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none"
                      stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                      <path d="M21 15v4a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-4"></path>
                      <polyline points="17 8 12 3 7 8"></polyline>
                      <line x1="12" y1="3" x2="12" y2="15"></line>
                    </svg>
                    อัปโหลดโลโก้
                  </button>
                  <button v-if="editCompanyData.logoUrl" type="button" class="btn-remove-logo" @click="removeLogo">
                    <svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none"
                      stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                      <polyline points="3 6 5 6 21 6"></polyline>
                      <path d="M19 6v14a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2V6m3 0V4a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v2">
                      </path>
                    </svg>
                    ลบโลโก้
                  </button>
                </div>
              </div>
              <p class="helper-text-small">รองรับไฟล์ .jpg, .png, .svg ขนาดไม่เกิน 2MB</p>
            </div>

            <div class="form-group">
              <label>ชื่อบริษัท</label>
              <input type="text" v-model="editCompanyData.companyName" placeholder="กรอกชื่อบริษัท" />
            </div>

            <div class="form-group">
              <label>อีเมลบริษัท</label>
              <input type="email" v-model="editCompanyData.email" placeholder="company@example.com" />
            </div>

            <div class="form-group">
              <label>เบอร์โทรศัพท์บริษัท</label>
              <input type="tel" v-model="editCompanyData.phone" placeholder="0123456789" />
            </div>

            <div class="panel-actions">
              <button class="btn-cancel" @click="closeEditCompanyPanel">ยกเลิก</button>
              <button class="btn-save" @click="saveCompanyInfo">บันทึก</button>
            </div>
          </div>
        </div>
      </div>
    </section>
  </DashboardLayout>
</template>

<script setup>
import DashboardLayout from "../../components/DashboardLayout.vue";
import { ref, reactive } from "vue";
import TabNavigation from "../../components/TabNavigation.vue";

// Tabs
const tabs = ["ข้อมูลสถานประกอบการ"];
const activeTab = ref(tabs[0]);

// Company Data
const companyData = reactive({
  logoUrl: "https://i.pinimg.com/1200x/17/e7/6b/17e76b8984df16dd5a32d84ba1a4628c.jpg",
  companyName: "GOOGLE THAILAND CO., LTD.",
  email: "xxxxxxxxx@gmail.com",
  phone: "0123456789"
});

// Panel states
const showEditCompanyPanel = ref(false);
const logoInput = ref(null);

// Edit data
const editCompanyData = reactive({
  logoUrl: "",
  companyName: "",
  email: "",
  phone: ""
});

// Methods for Company Info
const openEditCompanyPanel = () => {
  editCompanyData.logoUrl = companyData.logoUrl;
  editCompanyData.companyName = companyData.companyName;
  editCompanyData.email = companyData.email;
  editCompanyData.phone = companyData.phone;
  showEditCompanyPanel.value = true;
};

const closeEditCompanyPanel = () => {
  showEditCompanyPanel.value = false;
};

const saveCompanyInfo = () => {
  companyData.logoUrl = editCompanyData.logoUrl;
  companyData.companyName = editCompanyData.companyName;
  companyData.email = editCompanyData.email;
  companyData.phone = editCompanyData.phone;
  closeEditCompanyPanel();
};

// Logo Upload Methods
const triggerLogoInput = () => {
  logoInput.value.click();
};

const handleLogoUpload = (event) => {
  const file = event.target.files[0];
  if (file) {
    // ตรวจสอบประเภทไฟล์
    const validTypes = ['image/jpeg', 'image/png', 'image/svg+xml'];
    if (!validTypes.includes(file.type)) {
      alert('กรุณาอัปโหลดไฟล์ .jpg, .png, หรือ .svg เท่านั้น');
      return;
    }
    
    // ตรวจสอบขนาดไฟล์ (2MB)
    if (file.size > 2 * 1024 * 1024) {
      alert('ขนาดไฟล์ต้องไม่เกิน 2MB');
      return;
    }

    // สร้าง URL สำหรับแสดงภาพ
    const reader = new FileReader();
    reader.onload = (e) => {
      editCompanyData.logoUrl = e.target.result;
    };
    reader.readAsDataURL(file);
  }
};

const removeLogo = () => {
  if (confirm("คุณต้องการลบโลโก้หรือไม่?")) {
    editCompanyData.logoUrl = "";
    if (logoInput.value) {
      logoInput.value.value = "";
    }
  }
};
</script>

<style scoped>
@import url("https://fonts.googleapis.com/css2?family=Kanit:wght@400;500;600;700&display=swap");

* {
  font-family: "Kanit", "Prompt", "Sarabun", "Noto Sans Thai", sans-serif;
  box-sizing: border-box;
}

.profile-section {
  padding: 0;
}

/* ========== APPLICATION CONTENT ========== */
.application-content {
  padding: 20px 0;
  display: flex;
  flex-direction: column;
  gap: 35px;
}

.section-wrapper {
  display: flex;
  flex-direction: column;
  gap: 16px;
}

.section-title-outside {
  font-size: 22px;
  font-weight: 700;
  color: #222;
  margin: 0;
  padding-left: 4px;
}

.section-card {
  background: #fff;
  border-radius: 16px;
  padding: 40px;
  box-shadow: 0 2px 12px rgba(3, 114, 102, 0.07);
  border: 1px solid #eef0ef;
  position: relative;
}

/* ========== COMPANY PROFILE CONTAINER ========== */
.company-profile-container {
  display: flex;
  gap: 35px;
  align-items: flex-start;
}

.company-logo-section {
  flex-shrink: 0;
}

.company-logo {
  width: 140px;
  height: 140px;
  border-radius: 50%;
  background: #f5f5f5;
  display: flex;
  align-items: center;
  justify-content: center;
  overflow: hidden;
  border: 2px solid #e5e5e5;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.08);
}

.company-logo img {
  width: 100%;
  height: 100%;
  object-fit: cover;
  padding: 0px;
}

.company-logo svg {
  color: #bbb;
}

.company-info-section {
  flex: 1;
  display: flex;
  flex-direction: column;
  gap: 18px;
}

.company-name {
  font-size: 24px;
  font-weight: 700;
  color: #222;
  margin: 0;
  line-height: 1.3;
}

.info-items {
  display: flex;
  flex-direction: column;
  gap: 10px;
  margin-bottom: 6px;
}

.info-item {
  display: flex;
  align-items: center;
  gap: 10px;
  font-size: 15px;
  color: #666;
  margin: 0;
}

.info-item svg {
  color: #999;
  flex-shrink: 0;
}

.info-item span {
  color: #555;
}

.btn-edit {
  padding: 11px 24px;
  border: none;
  border-radius: 10px;
  font-size: 15px;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.2s;
  display: inline-block;
  background: #f0f8f7;
  color: #037266;
  border: 2px solid #cee5e2;
  align-self: flex-start;
}

.btn-edit:hover {
  background: #e0f2f0;
  border-color: #037266;
  transform: translateY(-1px);
  box-shadow: 0 4px 12px rgba(3, 114, 102, 0.15);
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
  width: 600px;
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
  background: #f0f0f0;
  color: #666;
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
  background: #e0e0e0;
  transform: scale(1.05);
}

.side-panel-inner {
  padding: 46px 34px 38px 34px;
  display: flex;
  flex-direction: column;
  gap: 24px;
}

.side-panel-title {
  font-size: 24px;
  font-weight: 700;
  color: #222;
  margin: 0;
}

.form-group {
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.form-group label {
  font-size: 14px;
  font-weight: 600;
  color: #555;
}

.form-group input {
  padding: 12px 16px;
  border: 1px solid #e0e0e0;
  border-radius: 10px;
  font-size: 15px;
  transition: all 0.2s;
  outline: none;
  font-family: "Kanit", "Prompt", "Sarabun", "Noto Sans Thai", sans-serif;
}

.form-group input:focus {
  border-color: #037266;
  box-shadow: 0 0 0 3px rgba(3, 114, 102, 0.1);
}

/* ========== LOGO UPLOAD SECTION ========== */
.logo-upload-section {
  display: flex;
  align-items: center;
  gap: 20px;
  padding: 20px;
  background: #fafafa;
  border: 1px solid #e5e5e5;
  border-radius: 12px;
}

.current-logo {
  width: 100px;
  height: 100px;
  border-radius: 50%;
  background: #fff;
  display: flex;
  align-items: center;
  justify-content: center;
  overflow: hidden;
  border: 2px solid #e0e0e0;
  flex-shrink: 0;
}

.current-logo img {
  width: 100%;
  height: 100%;
  object-fit: cover;
  padding: 0px;
}

.no-logo {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 100%;
  height: 100%;
  color: #ccc;
}

.logo-actions {
  display: flex;
  flex-direction: column;
  gap: 10px;
  flex: 1;
}

.btn-upload-logo,
.btn-remove-logo {
  padding: 10px 20px;
  border: none;
  border-radius: 8px;
  font-size: 14px;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.2s;
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 8px;
}

.btn-upload-logo {
  background: #f0f8f7;
  color: #037266;
  border: 2px solid #cee5e2;
}

.btn-upload-logo:hover {
  background: #e0f2f0;
  border-color: #037266;
  transform: translateY(-1px);
}

.btn-remove-logo {
  background: #fee;
  color: #e74c3c;
  border: 2px solid #fdd;
}

.btn-remove-logo:hover {
  background: #fdd;
  border-color: #e74c3c;
  transform: translateY(-1px);
}

.helper-text-small {
  font-size: 13px;
  color: #999;
  margin: 6px 0 0 0;
  line-height: 1.4;
}

.panel-actions {
  display: flex;
  justify-content: flex-end;
  gap: 12px;
  margin-top: 30px;
  padding-top: 25px;
  border-top: 1px solid #eee;
}

.btn-cancel,
.btn-save {
  padding: 12px 28px;
  border: none;
  border-radius: 10px;
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
  color: #333;
}

.btn-save {
  background: linear-gradient(135deg, #037266 0%, #026b5b 100%);
  color: #fff;
  box-shadow: 0 4px 12px rgba(3, 114, 102, 0.2);
}

.btn-save:hover {
  background: linear-gradient(135deg, #026b5b 0%, #025a4d 100%);
  transform: translateY(-1px);
  box-shadow: 0 6px 16px rgba(3, 114, 102, 0.3);
}

/* ========== RESPONSIVE ========== */
@media (max-width: 900px) {
  .company-profile-container {
    flex-direction: column;
    align-items: center;
    text-align: center;
  }

  .company-info-section {
    align-items: center;
  }

  .btn-edit {
    align-self: center;
  }

  .side-panel {
    width: 100%;
    max-width: 100%;
    border-radius: 0;
  }

  .side-panel-inner {
    padding: 30px 20px;
  }
}

@media (max-width: 650px) {
  .section-card {
    padding: 25px 20px;
  }

  .section-title-outside {
    font-size: 20px;
  }

  .company-profile-container {
    gap: 25px;
  }

  .company-logo {
    width: 120px;
    height: 120px;
  }

  .company-name {
    font-size: 20px;
  }

  .logo-upload-section {
    flex-direction: column;
    text-align: center;
  }

  .logo-actions {
    width: 100%;
  }

  .panel-actions {
    flex-direction: column-reverse;
  }

  .btn-cancel,
  .btn-save {
    width: 100%;
  }

  .side-close {
    top: 12px;
    right: 12px;
    width: 32px;
    height: 32px;
    font-size: 18px;
  }

  .side-panel-inner {
    padding: 20px 16px;
  }

  .side-panel-title {
    font-size: 20px;
  }
}
</style>