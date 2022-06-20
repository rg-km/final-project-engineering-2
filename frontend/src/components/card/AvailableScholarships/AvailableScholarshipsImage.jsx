import React from "react";

function AvailableScholarshipsImage({imageUrl}) {
    return (
        <div className="available-scholarships__image">
            <img src={imageUrl} alt="available scholarships"/>
        </div>
    )
}

export default AvailableScholarshipsImage;