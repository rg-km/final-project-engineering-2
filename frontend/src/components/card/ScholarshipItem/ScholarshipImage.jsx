import React from "react";

function ScholarshipsImage({imageUrl}) {
    return (
        <div className="scholarships__image">
            <img src={imageUrl} alt="scholarships"/>
        </div>
    )
}

export default ScholarshipsImage;