import React from "react";

function ScholarshipsBody({ name, faculty, level, completionTime }) {
  return (
    <div className="scholarships__body">
      <h3 className="scholarships__university-name">{name}</h3>
      <h4 className="scholarships__university-faculty">{faculty}</h4>
      <h5 className="scholarships__university-level">{level}</h5>
      <h5 className="scholarships__university-completion-time">
        {completionTime}
      </h5>
    </div>
  );
}

export default ScholarshipsBody;