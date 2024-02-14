
namespace TemplateBGMT
{
    public partial class App : Application
    {
        public App()
        {
            InitializeComponent();
            MainPage = new AppShell();
        }

        protected override Window CreateWindow(IActivationState? activationState)
        {
            Window window = base.CreateWindow(activationState); ;
            const int startWidth = 1000;
            const int startHeight = 1000;

            window.Width = startWidth;
            window.Height = startHeight;

            return window;
        }
    }
}
