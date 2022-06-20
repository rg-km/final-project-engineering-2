import React from "react";
import AvailableScholarshipsImage from "./AvailableScholarshipsImage";
import AvailableScholarshipsBody from "./AvailableSholarshipsBody";

function AvailableScholarshipsItem ({imageUrl, name, faculty, level, completionTime}) {
    return (
        <div className="available-scholarships-item">
            <AvailableScholarshipsImage imageUrl={imageUrl}/>
            <AvailableScholarshipsBody name={name} faculty={faculty} level={level} completionTime={completionTime}/>
        </div>
    )
}

export default AvailableScholarshipsItem;