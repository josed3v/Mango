
  $(document).ready(function () {
    // Fecha o menu ao clicar em um item
    $('.navbar-nav a').on('click', function () {
      $('.navbar-collapse').collapse('hide');
    });

    // Fecha o menu ao clicar no botão de toggle
    $('#navbarSideCollapse').on('click', function () {
      $('.navbar-collapse').collapse('hide');
    });
  });