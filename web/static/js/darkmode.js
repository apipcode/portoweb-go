/**
 * DARKMODE.JS — Toggle dark mode
 * Menyimpan preferensi dark mode di localStorage
 * Seperti tombol lampu baca di malam hari 🌙
 */

(function () {
    'use strict';

    // Ambil elemen toggle
    var toggle = document.getElementById('darkmode-toggle');
    if (!toggle) return;

    // Key localStorage untuk menyimpan preferensi
    var STORAGE_KEY = 'portfolio-darkmode';

    /**
     * setDarkMode mengaktifkan atau menonaktifkan dark mode
     * @param {boolean} isDark - true untuk dark mode, false untuk light mode
     */
    function setDarkMode(isDark) {
        if (isDark) {
            document.documentElement.setAttribute('data-theme', 'dark');
            toggle.textContent = '☀️';
            toggle.title = 'Mode terang';
        } else {
            document.documentElement.removeAttribute('data-theme');
            toggle.textContent = '🌙';
            toggle.title = 'Mode gelap';
        }

        // Simpan preferensi ke localStorage
        try {
            localStorage.setItem(STORAGE_KEY, isDark ? 'dark' : 'light');
        } catch (e) {
            // localStorage tidak tersedia (private browsing, dll)
        }
    }

    /**
     * loadPreference membaca preferensi dark mode yang tersimpan
     * Jika belum ada preferensi, cek system preference
     */
    function loadPreference() {
        var stored = null;
        try {
            stored = localStorage.getItem(STORAGE_KEY);
        } catch (e) {
            // Abaikan error localStorage
        }

        if (stored === 'dark') {
            setDarkMode(true);
        } else if (stored === 'light') {
            setDarkMode(false);
        } else {
            // Belum ada preferensi — cek preferensi sistem
            var prefersDark = window.matchMedia && window.matchMedia('(prefers-color-scheme: dark)').matches;
            setDarkMode(prefersDark);
        }
    }

    // Toggle saat tombol diklik
    toggle.addEventListener('click', function () {
        var isDark = document.documentElement.getAttribute('data-theme') === 'dark';
        setDarkMode(!isDark);
    });

    // Muat preferensi saat halaman dimuat
    loadPreference();
})();
