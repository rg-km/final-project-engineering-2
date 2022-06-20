import React from "react";
import AvailableScholarshipsList from "./AvailableScholarshipsList";
import "./AvailableScholarships.css"

function AvailableScholarships ({scholarshipsList}){
    return (
        <div className="available-scholarships">
            <h2>Available Scholarships</h2>
            <AvailableScholarshipsList scholarshipsList={scholarshipsList}/>
        </div>
    )
}

export default AvailableScholarships;