/**
 * FLIPBOOK.JS — Mesin flip-book custom
 * Mengatur navigasi halaman, animasi flip, dan interaksi pengguna
 * Tidak menggunakan library eksternal — murni vanilla JS + CSS 3D transforms
 */

(function () {
    'use strict';

    // ============================================
    // STATE — Variabel status flip-book
    // ============================================

    // Ambil semua elemen halaman
    const pages = document.querySelectorAll('.page');
    const totalPages = pages.length;

    // Halaman aktif saat ini (0 = cover)
    let currentPage = 0;

    // Elemen navigasi
    const prevBtn = document.getElementById('prev-page');
    const nextBtn = document.getElementById('next-page');
    const indicator = document.getElementById('page-indicator');

    // Label untuk setiap halaman
    const pageLabels = [
        'Cover',
        'About Me',
        'Experience',
        'Projects',
        'Tech Stack',
        'Contact'
    ];

    // ============================================
    // NAVIGASI — Fungsi untuk membalik halaman
    // ============================================

    /**
     * flipToPage membalik halaman ke nomor yang ditentukan
     * Halaman sebelum targetPage akan dibalik (class 'flipped')
     * Halaman setelah targetPage akan tidak dibalik
     * @param {number} targetPage - Nomor halaman tujuan (0-indexed)
     */
    function flipToPage(targetPage) {
        // Batasi range halaman
        if (targetPage < 0 || targetPage >= totalPages) return;

        currentPage = targetPage;

        // Balik atau buka setiap halaman sesuai posisi
        pages.forEach(function (page, index) {
            if (index < currentPage) {
                // Halaman sebelum halaman aktif — sudah dibalik
                page.classList.add('flipped');
            } else {
                // Halaman aktif dan setelahnya — belum dibalik
                page.classList.remove('flipped');
            }
        });

        // Update UI navigasi
        updateNavigation();
    }

    /**
     * nextPage membalik ke halaman berikutnya
     */
    function nextPage() {
        if (currentPage < totalPages) {
            flipToPage(currentPage + 1);
        }
    }

    /**
     * prevPage kembali ke halaman sebelumnya
     */
    function prevPage() {
        if (currentPage > 0) {
            flipToPage(currentPage - 1);
        }
    }

    /**
     * updateNavigation memperbarui tampilan tombol navigasi dan indikator
     */
    function updateNavigation() {
        // Disable tombol prev jika di halaman pertama
        if (prevBtn) prevBtn.disabled = currentPage <= 0;
        // Disable tombol next jika di halaman terakhir
        if (nextBtn) nextBtn.disabled = currentPage >= totalPages;
        // Update label halaman
        if (indicator) {
            if (currentPage === 0) {
                indicator.textContent = pageLabels[0];
            } else if (currentPage >= totalPages) {
                indicator.textContent = 'Back Cover';
            } else {
                indicator.textContent = pageLabels[currentPage] || ('Halaman ' + currentPage);
            }
        }
    }

    // ============================================
    // EVENT LISTENERS — Interaksi pengguna
    // ============================================

    // Klik tombol navigasi
    if (prevBtn) prevBtn.addEventListener('click', prevPage);
    if (nextBtn) nextBtn.addEventListener('click', nextPage);

    // Klik halaman untuk membalik
    pages.forEach(function (page, index) {
        page.addEventListener('click', function (e) {
            // Jangan balik jika klik pada link, form, input, atau button
            if (e.target.closest('a, form, input, textarea, button, .project-card')) return;

            // Cek apakah di mode mobile (lebar < 768px) — jangan flip di mobile
            if (window.innerWidth < 768) return;

            // Tentukan arah balik berdasarkan posisi klik
            var rect = page.getBoundingClientRect();
            var clickX = e.clientX - rect.left;

            if (clickX > rect.width * 0.5) {
                // Klik di sisi kanan — balik ke halaman berikutnya
                nextPage();
            } else {
                // Klik di sisi kiri — kembali ke halaman sebelumnya
                prevPage();
            }
        });
    });

    // Navigasi keyboard (panah kiri/kanan)
    document.addEventListener('keydown', function (e) {
        // Jangan handle jika user sedang mengetik di input/textarea
        if (e.target.tagName === 'INPUT' || e.target.tagName === 'TEXTAREA') return;

        if (e.key === 'ArrowRight' || e.key === 'ArrowDown') {
            e.preventDefault();
            nextPage();
        } else if (e.key === 'ArrowLeft' || e.key === 'ArrowUp') {
            e.preventDefault();
            prevPage();
        }
    });

    // ============================================
    // TOUCH / SWIPE — Navigasi sentuh untuk tablet
    // ============================================

    var touchStartX = 0;
    var touchStartY = 0;
    var isSwiping = false;

    document.addEventListener('touchstart', function (e) {
        touchStartX = e.touches[0].clientX;
        touchStartY = e.touches[0].clientY;
        isSwiping = true;
    }, { passive: true });

    document.addEventListener('touchend', function (e) {
        if (!isSwiping) return;
        isSwiping = false;

        var touchEndX = e.changedTouches[0].clientX;
        var touchEndY = e.changedTouches[0].clientY;
        var diffX = touchStartX - touchEndX;
        var diffY = touchStartY - touchEndY;

        // Hanya proses swipe horizontal yang cukup jauh (min 50px)
        // dan lebih horizontal daripada vertikal
        if (Math.abs(diffX) > 50 && Math.abs(diffX) > Math.abs(diffY)) {
            // Hanya proses swipe di mode desktop (tablet landscape)
            if (window.innerWidth >= 768) {
                if (diffX > 0) {
                    nextPage(); // Swipe kiri = halaman berikutnya
                } else {
                    prevPage(); // Swipe kanan = halaman sebelumnya
                }
            }
        }
    }, { passive: true });

    // ============================================
    // === VIEW TOGGLE LOGIC ===
    const viewToggleBtn = document.getElementById('view-toggle');
    const body = document.body;
    let isScrollingView = false;

    if (viewToggleBtn) {
        viewToggleBtn.addEventListener('click', toggleView);
    }

    function toggleView() {
        isScrollingView = !isScrollingView;

        if (isScrollingView) {
            body.classList.add('view-scrolling');
            viewToggleBtn.textContent = '📖'; // Icon buku untuk kembali
            viewToggleBtn.title = 'Kembali ke Tampilan Buku';
            // Reset flip state visually if needed, but CSS handles transforms
        } else {
            body.classList.remove('view-scrolling');
            viewToggleBtn.textContent = '📱'; // Icon HP untuk scroll
            viewToggleBtn.title = 'Ubah ke Tampilan Scroll';
            // Restore flip state
            flipToPage(currentPage);
        }
    }

    // Handle initial state if needed (e.g. from URL hash or localStorage)
    // For now, default is Flip-Book as requested.

    // ============================================
    // MOBILE NAVIGATION — Navigasi di layar kecil
    // ============================================

    // Di mobile, halaman ditampilkan sebagai scroll vertikal
    // Navigasi mobile menggunakan anchor links
    var mobileNavItems = document.querySelectorAll('.mobile-nav-item');
    var sections = ['cover', 'about', 'experience', 'projects', 'techstack', 'contact'];

    mobileNavItems.forEach(function (item) {
        item.addEventListener('click', function (e) {
            e.preventDefault();
            var section = item.getAttribute('data-section');
            var sectionIndex = sections.indexOf(section);

            if (sectionIndex >= 0 && sectionIndex < pages.length) {
                // Di mobile, scroll ke halaman yang sesuai
                var targetPage = pages[sectionIndex];
                if (targetPage) {
                    // Cari page-front di dalam halaman
                    var frontSide = targetPage.querySelector('.page-front');
                    if (frontSide) {
                        frontSide.scrollIntoView({ behavior: 'smooth', block: 'start' });
                    } else {
                        targetPage.scrollIntoView({ behavior: 'smooth', block: 'start' });
                    }
                }

                // Update active state
                mobileNavItems.forEach(function (nav) { nav.classList.remove('active'); });
                item.classList.add('active');
            }
        });
    });

    // Update mobile nav active state saat scroll
    if (window.innerWidth < 768) {
        var observerOptions = {
            root: null,
            threshold: 0.3
        };

        var observer = new IntersectionObserver(function (entries) {
            entries.forEach(function (entry) {
                if (entry.isIntersecting) {
                    var pageIndex = parseInt(entry.target.getAttribute('data-page'));
                    if (pageIndex >= 0 && pageIndex < mobileNavItems.length) {
                        mobileNavItems.forEach(function (nav) { nav.classList.remove('active'); });
                        mobileNavItems[pageIndex].classList.add('active');
                    }
                }
            });
        }, observerOptions);

        pages.forEach(function (page) {
            observer.observe(page);
        });
    }

    // ============================================
    // INISIALISASI
    // ============================================

    // Set state awal — mulai dari cover
    updateNavigation();
})();
