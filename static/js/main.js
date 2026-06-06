document.addEventListener('DOMContentLoaded', function() {
    var forms = document.querySelectorAll('form');
    forms.forEach(function(form) {
        var deleteBtn = form.querySelector('button[type="submit"].btn-danger');
        if (deleteBtn) {
            form.addEventListener('submit', function(e) {
                if (!confirm('确认删除此记录？此操作不可撤销。')) {
                    e.preventDefault();
                }
            });
        }
    });

    var currentPath = window.location.pathname;
    var navItems = document.querySelectorAll('.nav-item');
    navItems.forEach(function(item) {
        var href = item.getAttribute('href');
        if (href && currentPath.startsWith(href) && href !== '/') {
            item.classList.add('active');
        } else if (href === '/' && currentPath === '/') {
            item.classList.add('active');
        }
    });

    var alerts = document.querySelectorAll('.alert');
    alerts.forEach(function(alert) {
        setTimeout(function() {
            alert.style.transition = 'opacity 0.3s';
            alert.style.opacity = '0';
            setTimeout(function() { alert.remove(); }, 300);
        }, 5000);
    });

    var phoneInputs = document.querySelectorAll('input[name="phone"]');
    phoneInputs.forEach(function(input) {
        input.addEventListener('input', function() {
            this.value = this.value.replace(/[^\d+\-]/g, '');
        });
    });

    var ageInputs = document.querySelectorAll('input[name="age"]');
    ageInputs.forEach(function(input) {
        input.addEventListener('input', function() {
            if (this.value && parseInt(this.value) < 1) this.value = 1;
            if (this.value && parseInt(this.value) > 120) this.value = 120;
        });
    });
});
