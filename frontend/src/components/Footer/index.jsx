import React from "react";
import styles from "./Footer.module.css";

function Footer() {
    return (
        <div className={styles.footer__container}>
            <p className={styles.footer__title}>&copy; {new Date().getFullYear()} Raih Beasiswa | MSIB Kampus Merdeka x Ruangguru Web Development Final Project</p>
        </div>
    )
}

export default Footer;