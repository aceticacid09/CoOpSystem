<template>
  <div>
    <!-- ========== HEADER ========== -->
    <Header />

    <!-- ========== BANNER SECTION ========== -->
    <section class="banner">
      <div class="container">
        <div class="department-info">
          <div>
            <strong>{{ department.name }}</strong> <br>
            <small>{{ department.university }}</small>
          </div>
        </div>
      </div>
    </section>

    <!-- ========== NEWS & EVENTS SECTION ========== -->
    <NewsSection :maxHeight="'550px'" />

    <!-- ========== CO-OP SECTION ========== -->
    <section class="coop-section">
      <div class="coop-container">
        <h2 class="coop-title">สหกิจศึกษา</h2>
        <p class="coop-desc">
          รูปแบบการเรียนรู้ที่ผสานการเรียนในห้องกับการทำงานจริง
          นักศึกษาจะได้ออกไปปฏิบัติงานในสถานประกอบการที่เกี่ยวข้องกับสาขาวิชา
          เพื่อเสริมทักษะวิชาชีพ ประสบการณ์ทำงาน และเตรียมความพร้อมสู่ตลาดแรงงาน
          เป็นส่วนสำคัญของการพัฒนาศักยภาพนักศึกษาคณะวิทยาศาสตร์
        </p>
      </div>
    </section>

    <!-- ========== COORDINATORS SECTION ========== -->
    <section class="coordinators-section">
      <h2 class="coordinators-title">อาจารย์ผู้ประสานงานสหกิจศึกษา</h2>

      <swiper :slides-per-view="5" :centered-slides="true" :space-between="30" :loop="true" :effect="'coverflow'"
        :coverflow-effect="{
          rotate: 0,
          stretch: 0,
          depth: 150,
          modifier: 2,
          slideShadows: false,
        }" :slide-to-clicked-slide="true" class="coordinators-swiper" @slideChange="onSlideChange">
        <swiper-slide v-for="(teacher, index) in teachers" :key="index">
          <div class="teacher-card-wrapper">
            <div class="teacher-card" :class="{ active: activeIndex === index }">
              <img :src="teacher.image" :alt="teacher.name" class="teacher-img" />
            </div>
          </div>
        </swiper-slide>
      </swiper>

      <div class="teacher-info" v-if="teachers[activeIndex]">
        <h3 class="teacher-name">{{ teachers[activeIndex].name }}</h3>
        <p class="teacher-position">{{ teachers[activeIndex].position }}</p>
      </div>
    </section>

    <!-- ========== Back to Top ========== -->
    <BackToTop />

    <!-- ========== FOOTER ========== -->
    <Footer />
  </div>
</template>

<script setup>
import { reactive, ref, onMounted, onUnmounted } from "vue";
import { Swiper, SwiperSlide } from "swiper/vue";
import "swiper/css";
import "swiper/css/effect-coverflow";
import Header from "../components/Header.vue";
import Footer from "../components/Footer.vue";
import NewsSection from "../components/NewsSections.vue";
import BackToTop from "../components/BackToTop.vue";

/* ========================================
   DEPARTMENT INFORMATION
======================================== */
const department = reactive({
  name: "สหกิจศึกษา",
  university: "คณะวิทยาศาสตร์ มหาวิทยาลัยศิลปากร",
});

/* ========================================
   TEACHERS / COORDINATORS DATA & STATE
======================================== */
const activeIndex = ref(0);

const teachers = ref([
  {
    name: "ผศ.ดร.สัจอาภรณ์ ไววรรยา",
    position: "ภาควิชาคอมพิวเตอร์",
    image: "https://www.cp.su.ac.th/image/crop/856",
  },
  {
    name: "อ. สมชาย ใจดี",
    position: "ภาควิชาคณิตศาสตร์",
    image:
      "https://i.pinimg.com/736x/63/53/d9/6353d9fff14cc31af369dd0254fd8c97.jpg",
  },
  {
    name: "อ. สมหญิง เก่งงาน",
    position: "ภาควิชาเคมี",
    image:
      "https://i.pinimg.com/736x/63/53/d9/6353d9fff14cc31af369dd0254fd8c97.jpg",
  },
  {
    name: "อ. ธนพล วิริยะ",
    position: "ภาควิชาชีววิทยา",
    image:
      "https://i.pinimg.com/736x/63/53/d9/6353d9fff14cc31af369dd0254fd8c97.jpg",
  },
  {
    name: "อ. นฤมล ศรีสุข",
    position: "ภาควิชาฟิสิกส์",
    image:
      "https://i.pinimg.com/736x/63/53/d9/6353d9fff14cc31af369dd0254fd8c97.jpg",
  },
  {
    name: "อ. พรชัย อารีย์",
    position: "ภาควิชาวิทยาศาสตร์สิ่งแวดล้อม",
    image:
      "https://i.pinimg.com/736x/63/53/d9/6353d9fff14cc31af369dd0254fd8c97.jpg",
  },
  {
    name: "อ. สุรีย์รัตน์ วัฒนะ",
    position: "ภาควิชาสถิติ",
    image:
      "https://i.pinimg.com/736x/63/53/d9/6353d9fff14cc31af369dd0254fd8c97.jpg",
  },
  {
    name: "อ. กิตติศักดิ์ ชาญชัย",
    position: "ภาควิชาจุลชีววิทยา",
    image:
      "https://i.pinimg.com/736x/63/53/d9/6353d9fff14cc31af369dd0254fd8c97.jpg",
  }
]);

const onSlideChange = (swiper) => {
  activeIndex.value = swiper.realIndex;
};

</script>

<style scoped>
@import url("https://fonts.googleapis.com/css2?family=Kanit:wght@400;700&display=swap");

/* ========================================
   GLOBAL STYLES
======================================== */
* {
  box-sizing: border-box;
  font-family: "Kanit", "Prompt", "Sarabun", "Noto Sans Thai", sans-serif;
}

.container {
  max-width: 1200px;
  margin: auto;
  padding: 0 20px;
}

/* ========================================
   BANNER SECTION
======================================== */
.banner {
  position: relative;
  background: linear-gradient(to right,
      rgba(255, 255, 255, 0.9),
      rgba(80, 118, 118, 0.6)),
    url("https://www.su.ac.th/th/images-faculty/07_fac_sci_01.jpg") center bottom/cover no-repeat;
  padding: 80px 30px;
  color: #000;
  text-align: left;
  margin-bottom: 60px;
}

.department-info {
  display: flex;
  flex-direction: column;
  align-items: flex-start;
  gap: 10px;
  margin-left: 20px;
}

.department-info strong {
  font-size: 44px;
  color: #037266;
}

.department-info small {
  font-size: 18px;
  color: #000000;
}

/* ========================================
   CO-OP SECTION
======================================== */
.coop-section {
  background: linear-gradient(to right, #e6f2f1, #fff7e6);
  padding: 80px 20px;
  text-align: center;
  margin-top: 60px;
}

.coop-container {
  max-width: 900px;
  margin: auto;
}

.coop-title {
  font-size: 30px;
  font-weight: 700;
  color: #037266;
  margin-bottom: 20px;
}

.coop-desc {
  font-size: 16px;
  color: #444;
  line-height: 1.8;
  max-width: 800px;
  margin: 0 auto;
}

/* ========================================
   COORDINATORS SECTION
======================================== */
.coordinators-section {
  text-align: center;
  padding: 80px 20px;
  background: #fff;
}

.coordinators-title {
  font-size: 28px;
  font-weight: 700;
  color: #037266;
  margin-bottom: 60px;
}

.coordinators-swiper {
  max-width: 1000px;
  margin: auto;
  padding: 20px 0;
}

.teacher-card-wrapper {
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 10px;
}

.teacher-card {
  width: 150px;
  height: 150px;
  border-radius: 50%;
  overflow: hidden;
  opacity: 0.5;
  transition: all 0.3s ease;
  position: relative;
  background: #f5f5f5;
}

.teacher-card.active {
  transform: scale(1.3);
  opacity: 1;
  border: 4px solid #c0c0c0;
}

.teacher-img {
  width: 100%;
  height: 100%;
  object-fit: cover;
  object-position: center;
  display: block;
  position: absolute;
  top: 0;
  left: 0;
}

.teacher-info {
  margin-top: 20px;
}

.teacher-name {
  font-size: 18px;
  font-weight: 600;
  color: #037266;
}

.teacher-position {
  font-size: 14px;
  color: #555;
}

/* ========================================
   RESPONSIVE DESIGN
======================================== */
@media (max-width: 768px) {
  .news-card {
    flex-direction: column;
  }

  .news-img img {
    width: 100%;
    height: auto;
  }

  .side-images img {
    width: 100%;
    height: auto;
  }
}
</style>