// frontend/src/data/applicationData.js

export const mockApplicationData = {
    applicant: {
        fullName: "สมชาย ใจดี",
        email: "somchai.jaidee@gmail.com",
        phone: "0812345678"
    },

    // ✅ แก้ไขให้เป็น array
    documents: {
        resumes: [
            { id: 1, name: "Somchai_Resume_2024.pdf" },
            { id: 2, name: "Somchai_Resume_English.pdf" },
            { id: 3, name: "Somchai_CV_Professional.pdf" }
        ],
        coverLetterName: "Somchai_CoverLetter.pdf",
        transcriptName: "Transcript_Somchai.pdf"
    },

    applicationDate: "15 มกราคม 2568"
};

export const formatThaiDate = (dateStr) => {
    const monthsThai = [
        "มกราคม", "กุมภาพันธ์", "มีนาคม", "เมษายน", "พฤษภาคม", "มิถุนายน",
        "กรกฎาคม", "สิงหาคม", "กันยายน", "ตุลาคม", "พฤศจิกายน", "ธันวาคม"
    ];

    if (dateStr.includes('/')) {
        const parts = dateStr.split("/");
        const day = parts[0];
        const month = parseInt(parts[1]) - 1;
        const year = parseInt(parts[2]) + 2500;
        return `${day} ${monthsThai[month]} ${year}`;
    }

    return dateStr;
};

export const getApplicationByJobId = (jobId) => {
    return {
        ...mockApplicationData,
        jobId: jobId
    };
};

export default mockApplicationData;