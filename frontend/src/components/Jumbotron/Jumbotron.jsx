import React from "react";
import styles from "./Jumbotron.module.css";

function Jumbotron() {
  return (
    <div className={styles.jumbotron__container}>
      <div class={styles.jumbotron__inner}>
        <h2 class={styles.jumbotron__title}>
          Raih Cita-cita dengan Raih Beasiswa
        </h2>
        <p class={styles.jumbotron__tagline}>
          Memajukan dan Memeratakan Pendidikan Untuk Capai
          Tujuan Pendidikan Nasional Indonesia
        </p>
      </div>
    </div>
  );
}

export default Jumbotron;
