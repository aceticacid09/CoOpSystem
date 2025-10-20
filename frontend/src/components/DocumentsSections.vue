<!-- frontend/src/components/DocumentsSections.vue -->
<template>
  <section class="documents-section container">
    <!-- Tabs -->
    <TabNavigation v-model="activeTab" :tabs="tabs" />


    <!-- Accordion -->
    <div class="accordion">
      <div v-for="(dept, index) in filteredDepartments" :key="index" class="accordion-item">
        <!-- Header -->
        <div class="accordion-header" @click="toggleAccordion(index)">
          <span>{{ dept.name }}</span>
          <svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" fill="none" stroke="currentColor"
            stroke-width="2" stroke-linecap="round" stroke-linejoin="round" :class="{ rotated: activeIndex === index }">
            <polyline points="6 9 12 15 18 9" />
          </svg>
        </div>

        <!-- Content -->
        <transition name="accordion">
          <div v-if="activeIndex === index" class="accordion-content">
            <ul>
              <li v-for="(doc, i) in dept.documents" :key="i">
                <a :href="doc.link" target="_blank">
                  {{ doc.code }} - {{ doc.title }}
                </a>
              </li>
            </ul>
          </div>
        </transition>
      </div>
    </div>
  </section>
</template>

<script setup>
import { ref, computed } from "vue";
import TabNavigation from "../components/TabNavigation.vue";

// แท็บเมนู
const tabs = ["แบบฟอร์มทั่วไป", "คู่มือ"];
const activeTab = ref(tabs[0]);

// เก็บ index accordion ที่เปิดอยู่
const activeIndex = ref(null);

// toggle เปิด-ปิด accordion
const toggleAccordion = (index) => {
  activeIndex.value = activeIndex.value === index ? null : index;
};

const departments = [
  // ===============================
  // แบบฟอร์มทั่วไป
  // ===============================
  {
    name: "ภาควิชาคณิตศาสตร์",
    type: "แบบฟอร์มทั่วไป",
    documents: [
      { code: "MATH-01", title: "คำร้องทั่วไป (คณิตศาสตร์)", link: "#" },
      { code: "MATH-02", title: "แบบฟอร์มสมัครเข้าร่วมโครงการสหกิจศึกษา", link: "#" },
      { code: "MATH-03", title: "แบบฟอร์มแจ้งผลการปฏิบัติงาน", link: "#" },
    ],
  },
  {
    name: "ภาควิชาเคมี",
    type: "แบบฟอร์มทั่วไป",
    documents: [
      { code: "CHEM-01", title: "คำร้องทั่วไป (เคมี)", link: "#" },
      { code: "CHEM-02", title: "ใบสมัครโครงการสหกิจศึกษา", link: "#" },
      { code: "CHEM-03", title: "แบบฟอร์มรายงานความก้าวหน้า", link: "#" },
    ],
  },
  {
    name: "ภาควิชาชีววิทยา",
    type: "แบบฟอร์มทั่วไป",
    documents: [
      { code: "BIO-01", title: "คำร้องทั่วไป (ชีววิทยา)", link: "#" },
      { code: "BIO-02", title: "ใบสมัครโครงการสหกิจศึกษา", link: "#" },
    ],
  },
  {
    name: "ภาควิชาฟิสิกส์",
    type: "แบบฟอร์มทั่วไป",
    documents: [
      { code: "PHYS-01", title: "คำร้องทั่วไป (ฟิสิกส์)", link: "#" },
      { code: "PHYS-02", title: "แบบฟอร์มลงทะเบียนสหกิจศึกษา", link: "#" },
    ],
  },
  {
    name: "ภาควิชาวิทยาศาสตร์สิ่งแวดล้อม",
    type: "แบบฟอร์มทั่วไป",
    documents: [
      { code: "ENV-01", title: "คำร้องทั่วไป (วิทยาศาสตร์สิ่งแวดล้อม)", link: "#" },
      { code: "ENV-02", title: "ใบสมัครเข้าร่วมโครงการสหกิจศึกษา", link: "#" },
    ],
  },
  {
    name: "ภาควิชาสถิติ",
    type: "แบบฟอร์มทั่วไป",
    documents: [
      { code: "STAT-01", title: "คำร้องทั่วไป (สถิติ)", link: "#" },
      { code: "STAT-02", title: "แบบฟอร์มยืนยันสถานประกอบการ", link: "#" },
    ],
  },
  {
    name: "ภาควิชาคอมพิวเตอร์",
    type: "แบบฟอร์มทั่วไป",
    documents: [
      { code: "COMP-01", title: "คำร้องทั่วไป (คอมพิวเตอร์)", link: "#" },
      { code: "COMP-02", title: "ใบสมัครเข้าร่วมโครงการสหกิจศึกษา", link: "#" },
      { code: "COMP-03", title: "แบบฟอร์มส่งรายงานประจำสัปดาห์", link: "#" },
    ],
  },
  {
    name: "ภาควิชาจุลชีววิทยา",
    type: "แบบฟอร์มทั่วไป",
    documents: [
      { code: "MICRO-01", title: "คำร้องทั่วไป (จุลชีววิทยา)", link: "#" },
      { code: "MICRO-02", title: "ใบสมัครโครงการสหกิจศึกษา", link: "#" },
    ],
  },
  {
    name: "ภาควิชาคณิตศาสตร์ประยุกต์",
    type: "แบบฟอร์มทั่วไป",
    documents: [
      { code: "APPMATH-01", title: "คำร้องทั่วไป (คณิตศาสตร์ประยุกต์)", link: "#" },
      { code: "APPMATH-02", title: "ใบสมัครโครงการสหกิจศึกษา", link: "#" },
    ],
  },

  // ===============================
  // คู่มือ
  // ===============================
  {
    name: "คู่มือนักศึกษา",
    type: "คู่มือ",
    documents: [
      { code: "GUIDE-STUDENT", title: "คู่มือนักศึกษาโครงการสหกิจศึกษา", link: "#" },
    ],
  },
  {
    name: "คู่มือสถานประกอบการ",
    type: "คู่มือ",
    documents: [
      { code: "GUIDE-EMPLOYER", title: "คู่มือสำหรับสถานประกอบการ", link: "#" },
    ],
  },
  {
    name: "คู่มืออาจารย์",
    type: "คู่มือ",
    documents: [
      { code: "GUIDE-TEACHER", title: "คู่มือสำหรับอาจารย์ที่ปรึกษา", link: "#" },
    ],
  },
];

// กรองตามแท็บ
const filteredDepartments = computed(() =>
  departments.filter((d) => d.type === activeTab.value)
);
</script>

<style scoped>
/* Global */
* {
  font-family: "Kanit", "Prompt", "Sarabun", "Noto Sans Thai", sans-serif;
  box-sizing: border-box;
}

/* Section */
.documents-section {
  padding: 0;
}

/* Tabs */
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
  transition: 0.2s;
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

/* Accordion */
.accordion {
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.accordion-item {
  border: 1px solid #ddd;
  border-radius: 10px;
  overflow: hidden;
  background: #fff;
  transition: box-shadow 0.2s;
}

.accordion-item:hover {
  box-shadow: 0 3px 10px rgba(0, 0, 0, 0.08);
}

.accordion-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 14px 18px;
  font-size: 15px;
  font-weight: 600;
  color: #333;
  cursor: pointer;
  background: #f9f9f9;
  transition: background 0.2s;
}

.accordion-header:hover {
  background: #f1f1f1;
}

.accordion-header svg {
  transition: transform 0.25s;
}

.accordion-header svg.rotated {
  transform: rotate(180deg);
}

.accordion-content {
  padding: 15px 20px;
  background: #fff;
  border-top: 1px solid #eee;
}

.accordion-content ul {
  margin: 0;
  padding: 0;
  list-style: none;
}

.accordion-content li {
  margin: 8px 0;
}

.accordion-content a {
  color: #037266;
  text-decoration: none;
  font-size: 14px;
  transition: color 0.2s;
}

.accordion-content a:hover {
  color: #025c51;
  text-decoration: underline;
}

/* Transition */
.accordion-enter-active,
.accordion-leave-active {
  transition: all 0.25s ease;
}

.accordion-enter-from,
.accordion-leave-to {
  max-height: 0;
  opacity: 0;
  padding: 0 20px;
}

.accordion-enter-to,
.accordion-leave-from {
  max-height: 400px;
  opacity: 1;
}
</style>
