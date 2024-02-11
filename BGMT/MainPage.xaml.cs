using CommunityToolkit.Maui.Storage;
using System.Runtime.CompilerServices;

namespace BGMT
{
    public partial class MainPage : ContentPage
    {
        private string projectFolderPath = "";
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
                if (!folder.IsSuccessful) throw new Exception("Failed Folder Selection");
                this.projectFolderPath = folder.Folder.Path;
                target.Text = this.projectFolderPath;

                this.menuItemImport.IsEnabled = true;
                this.menuItemExport.IsEnabled = true;
            }
            catch (Exception ex) { Console.WriteLine(ex.ToString()); }
        }
        private async void NavBarFile_Open(object sender, EventArgs e)
        {
            try
            {
                var result = await FilePicker.Default.PickAsync();
                if (result != null)
                {
                    if (result.FileName.EndsWith("txt", StringComparison.OrdinalIgnoreCase))
                    {
                        using var stream = await result.OpenReadAsync();
                        using var reader = new StreamReader(stream);
                        string text = await reader.ReadToEndAsync();

                        /* Uncomment the below and consider where the path should be. Not sure yet. */
                        //string outputPath = "path/to/output.txt";
                        // using var writer = new StreamWriter(outputPath);
                        //await writer.WriteAsync(text);
                    }
                    target.Text = result.FileName;
                }
            }
            catch (Exception ex) { Console.WriteLine(ex.ToString()); }
        }
        private async void NavBarFile_Import(object sender, EventArgs e)
        {
            try
            {
                var result = await FilePicker.Default.PickAsync();
                if (result != null)
                {
                    if (result.FileName.EndsWith(".d", StringComparison.OrdinalIgnoreCase) || result.FileName.EndsWith(".tp2", StringComparison.OrdinalIgnoreCase) || result.FileName.EndsWith(".baf", StringComparison.OrdinalIgnoreCase))
                    {
                        using var stream = await result.OpenReadAsync();
                        using var reader = new StreamReader(stream);
                        string text = await reader.ReadToEndAsync();

                        /* Uncomment the below and consider where the path should be. Not sure yet. */
                        //string outputPath = "path/to/output.txt";
                        // using var writer = new StreamWriter(outputPath);
                        //await writer.WriteAsync(text);
                    }
                    target.Text = result.FileName;
                }
            }
            catch (Exception ex) { Console.WriteLine(ex.ToString()); }
        }
        private void NavBarFile_Export(object sender, EventArgs e)
        {
            target.Text = "Export Pressed";
        }
        private async void NavBarFile_Exit(object sender, EventArgs e)
        {
            bool answer = await DisplayAlert("Alert", "Are you sure you want to exit?", "Exit", "Cancel");
            if (answer)
            {
                Application.Current?.Quit();
            }
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
