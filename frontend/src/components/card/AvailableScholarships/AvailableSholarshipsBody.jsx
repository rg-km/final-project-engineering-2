import React from "react";

function AvailableScholarshipsBody({ name, faculty, level, completionTime }) {
  return (
    <div className="available-scholarships__body">
      <h3 className="available-scholarships__university-name">{name}</h3>
      <h4 className="available-scholarships__university-faculty">{faculty}</h4>
      <h5 className="available-scholarships__university-level">{level}</h5>
      <h5 className="available-scholarships__university-completion-time">
        {completionTime}
      </h5>
    </div>
  );
}

export default AvailableScholarshipsBody;
