using CommunityToolkit.Maui.Storage;
using System.Runtime.CompilerServices;

namespace BGMT
{
    public partial class MainPage : ContentPage
    {
        private string projectFolderPath;
        public MainPage()
        {
            InitializeComponent();

        }

        //File Subitem Clicked Functionality
        private async void NavBarFile_New(object sender, EventArgs e)
        {
            try
            {
                var folder = await FolderPicker.PickAsync(default);
                this.projectFolderPath = folder.Folder.Path;
                target.Text = this.projectFolderPath;

                this.menuItemImport.IsEnabled = true;
                this.menuItemExport.IsEnabled = true;
            }
            catch (Exception ex) { Console.WriteLine(ex.ToString()); }
        }
        private void NavBarFile_Open(object sender, EventArgs e)
        {
            target.Text = "Open Pressed";

        }
        private void NavBarFile_Import(object sender, EventArgs e)
        {
            target.Text = "Import Pressed";
        }
        private void NavBarFile_Export(object sender, EventArgs e)
        {
            target.Text = "Export Pressed";
        }
        private async Task NavBarFile_Exit(object sender, EventArgs e)
        {
            bool answer = await DisplayAlert("Alert", "Are you sure you want to exit?", "Exit", "Cancel");
            if (answer)
            {
                Application.Current?.Quit();
            }
            target.Text = "Import Pressed";
            target.Text = "Exit Pressed";
        }

        //Help Subitem Clicked Functionality
        private void NavBarHelp_Settings(object sender, EventArgs e)
        {
            target.Text = "Settings Pressed";
        }
        private void NavBarHelp_Help(object sender, EventArgs e)
        {
            target.Text = "Help Pressed";
        }

        //Mod Element Buttons
        private void ModElement_AddDialogue(object sender, EventArgs e)
        {
            target.Text = "Dialogue Pressed";
        }

        private void ModElement_AddScript(object sender, EventArgs e)
        {
            target.Text = "Script Pressed";
        }

        private void ModElement_AddInstallation(object sender, EventArgs e)
        {
            target.Text = "Installation Pressed";
        }
    }

}
