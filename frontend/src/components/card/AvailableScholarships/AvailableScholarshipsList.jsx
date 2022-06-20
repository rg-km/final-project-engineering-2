import React from "react";
import AvailableScholarshipsItem from "./AvailableScholarshipsItem";

function AvailableScholarshipsList({ scholarshipsList }) {
  return (
    <div className="available-scholarships-list">
      {scholarshipsList.map((scholarshipsItem) => (
        <AvailableScholarshipsItem
          key={scholarshipsItem?.id}
          name={scholarshipsItem?.name}
          faculty={scholarshipsItem?.faculty}
          level={scholarshipsItem?.level}
          completionTime={scholarshipsItem?.completionTime}
          imageUrl={scholarshipsItem?.imageUrl}
        />
      ))}
    </div>
  );
}

export default AvailableScholarshipsList;
