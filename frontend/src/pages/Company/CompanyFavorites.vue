<!-- frontend/src/pages/Student/StudentJobs.vue -->
<template>
    <DashboardLayout role="company">
        <FavoriteSections :saved-news-ids="savedNewsIds" :saved-jobs-ids="savedJobsIds"
            @remove-news-bookmark="handleRemoveNewsBookmark" @remove-job-bookmark="handleRemoveJobBookmark" />
    </DashboardLayout>
</template>

<script setup>
import { ref } from 'vue';
import DashboardLayout from '../../components/DashboardLayout.vue';
import FavoriteSections from '../../components/FavoriteSections.vue';

const savedNewsIds = ref([2, 3, 5]);
const savedJobsIds = ref([1, 3, 5]);

// Handle remove bookmark
const handleRemoveNewsBookmark = (newsId) => {
    const index = savedNewsIds.value.indexOf(newsId);
    if (index > -1) {
        savedNewsIds.value.splice(index, 1);
    }
    console.log('Removed news bookmark:', newsId);
    // ในโปรเจ็กต์จริง: เรียก API เพื่ออัพเดทฐานข้อมูล
};

const handleRemoveJobBookmark = (jobId) => {
    const index = savedJobsIds.value.indexOf(jobId);
    if (index > -1) {
        savedJobsIds.value.splice(index, 1);
    }
    console.log('Removed job bookmark:', jobId);
    // ในโปรเจ็กต์จริง: เรียก API เพื่ออัพเดทฐานข้อมูล
};
</script>

<style scoped></style>