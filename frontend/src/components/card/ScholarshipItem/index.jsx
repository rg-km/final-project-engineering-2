import React from "react";
import ScholarshipsImage from "./ScholarshipImage";
import ScholarshipsBody from "./ScholarshipBody";

function ScholarshipsItem({
  id,
  imageUrl,
  name,
  faculty,
  level,
  completionTime,
  onClickScholarship,
}) {
  return (
    <div className="scholarships-item" onClick={onClickScholarship}>
      <ScholarshipsImage imageUrl={imageUrl} />
      <ScholarshipsBody
        name={name}
        faculty={faculty}
        level={level}
        completionTime={completionTime}
      />
    </div>
  );
}

export default ScholarshipsItem;
