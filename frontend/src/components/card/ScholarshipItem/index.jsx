import React from "react";
import { useNavigate } from "react-router-dom";
import ScholarshipsImage from "./ScholarshipImage";
import ScholarshipsBody from "./ScholarshipBody";


function ScholarshipsItem ({id, imageUrl, name, faculty, level, completionTime}) {
    const navigate = useNavigate();
    return (
        <div className="scholarships-item" onClick={() => {
            localStorage.setItem("id_beasiswa", id)
            console.log(localStorage.getItem("id_beasiswa"))
        }}>
            <ScholarshipsImage imageUrl={imageUrl}/>
            <ScholarshipsBody name={name} faculty={faculty} level={level} completionTime={completionTime}/>
        </div>
    )
}

export default ScholarshipsItem;