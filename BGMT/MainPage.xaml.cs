using System.Diagnostics;
using System.Runtime.CompilerServices;

namespace BGMT
{
    public partial class MainPage : ContentPage
    {

        public MainPage()
        {
            InitializeComponent();
        }

        // File Subitem Clicked Functionality
        void NavBarFile_New(object sender, EventArgs e) { target.Text = "New Pressed"; }
        void NavBarFile_Open(object sender, EventArgs e) { target.Text = "Open Pressed"; }
        void NavBarFile_Import(object sender, EventArgs e) { target.Text = "Import Pressed"; }
        void NavBarFile_Export(object sender, EventArgs e) { target.Text = "Export Pressed"; }
        async void NavBarFile_Exit(object sender, EventArgs e)
        {
            bool answer = await DisplayAlert("Alert", "Are you sure you want to exit?", "Exit", "Cancel");
            if (answer) 
            {
                Application.Current?.Quit();
            }
        }

        // Help Subitem Clicked Functionality
        void NavBarHelp_Settings(object sender, EventArgs e) { target.Text = "Settings Pressed"; }
        void NavBarHelp_Help(object sender, EventArgs e) { target.Text = "Help Pressed"; }
    }

}
