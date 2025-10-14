<!-- frontend/src/pages/Student/StudentJobs.vue -->
<template>
  <DashboardLayout role="company">
    <JobsSections 
      user-role="company"
      :show-bookmark="true"
      :show-actions="false"
      :initial-bookmarked="[1, 3, 5]"
      :initial-saved="[2, 4]"
      :initial-applied=null
    />
  </DashboardLayout>
</template>

<script setup>
import DashboardLayout from '../../components/DashboardLayout.vue';
import JobsSections from "../../components/JobsSections.vue";
</script>

<style scoped></style>