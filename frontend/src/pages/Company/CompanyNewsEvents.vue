<!-- frontend/src/pages/Company/CompanyNewsEvents.vue -->
<template>
  <DashboardLayout role="company">
    <NewsSections 
      :show-bookmark="true"
      :initial-bookmarked="[1, 2, 4]"
      :initial-saved="[3, 5]"
    />
  </DashboardLayout>
</template>

<script setup>
import DashboardLayout from '../../components/DashboardLayout.vue';
import NewsSections from "../../components/NewsSections.vue";
</script>

<style scoped></style>
