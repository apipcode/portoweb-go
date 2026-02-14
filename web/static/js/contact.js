/**
 * CONTACT.JS — Validasi dan submit form kontak
 * Mengirim data form ke backend via fetch API
 * Menampilkan feedback langsung di halaman tanpa reload
 */

(function () {
    'use strict';

    // Ambil elemen form dan feedback
    var form = document.getElementById('contact-form');
    var feedback = document.getElementById('form-feedback');

    if (!form || !feedback) return;

    /**
     * showFeedback menampilkan pesan feedback ke pengguna
     * @param {string} message - Pesan yang ditampilkan
     * @param {string} type - Tipe pesan: 'success' atau 'error'
     */
    function showFeedback(message, type) {
        feedback.textContent = message;
        feedback.className = 'form-feedback ' + type;
    }

    /**
     * clearFeedback menghapus pesan feedback
     */
    function clearFeedback() {
        feedback.textContent = '';
        feedback.className = 'form-feedback';
    }

    /**
     * validateForm melakukan validasi client-side sebelum submit
     * @returns {object|null} - Object form data jika valid, null jika tidak
     */
    function validateForm() {
        var name = form.querySelector('[name="name"]').value.trim();
        var email = form.querySelector('[name="email"]').value.trim();
        var message = form.querySelector('[name="message"]').value.trim();

        // Validasi nama — minimal 2 karakter
        if (name.length < 2) {
            showFeedback('Nama minimal 2 karakter ya!', 'error');
            return null;
        }

        // Validasi email — format dasar
        var emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
        if (!emailRegex.test(email)) {
            showFeedback('Format email tidak valid.', 'error');
            return null;
        }

        // Validasi pesan — minimal 10 karakter
        if (message.length < 10) {
            showFeedback('Pesan minimal 10 karakter. Cerita lebih banyak dong!', 'error');
            return null;
        }

        // Validasi pesan — maksimal 2000 karakter
        if (message.length > 2000) {
            showFeedback('Pesan terlalu panjang (maks 2000 karakter).', 'error');
            return null;
        }

        return { name: name, email: email, message: message };
    }

    // Handle submit form
    form.addEventListener('submit', function (e) {
        e.preventDefault();
        clearFeedback();

        // Validasi client-side
        var data = validateForm();
        if (!data) return;

        // Tampilkan loading state
        var submitBtn = form.querySelector('.submit-btn');
        var originalText = submitBtn.textContent;
        submitBtn.textContent = 'Mengirim...';
        submitBtn.disabled = true;

        // Kirim data ke backend via fetch API
        fetch('/api/contact', {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(data)
        })
            .then(function (response) { return response.json(); })
            .then(function (result) {
                if (result.success) {
                    showFeedback(result.message || 'Pesan berhasil dikirim!', 'success');
                    form.reset(); // Kosongkan form
                } else {
                    showFeedback(result.message || 'Gagal mengirim pesan.', 'error');
                }
            })
            .catch(function (err) {
                showFeedback('Terjadi kesalahan jaringan. Coba lagi nanti.', 'error');
                console.error('Contact form error:', err);
            })
            .finally(function () {
                // Kembalikan tombol ke state normal
                submitBtn.textContent = originalText;
                submitBtn.disabled = false;
            });
    });
})();
